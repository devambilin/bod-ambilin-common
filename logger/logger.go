package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Pre-defined formats
const (
	DefaultRequestSize = 1000000 // default chan queue size in async logging

	//DefaultFormat     = "%-23s [%5s] %-20s - %s"
	DefaultFormat     = "%-23s [%5s] - %s"
	DefaultTimeFormat = "2006-01-02 15:04:05.000"
)

// Log Level
const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

type request struct {
	time   time.Time
	level  string
	name   string
	format string
	v      []interface{}
}

// SimpleLogger create simple logger
func SimpleLogger() *Logger {
	return CustomLogger("", false)
}

// CustomLogger create customized logger
func CustomLogger(file string, sync bool) *Logger {

	logger := new(Logger)
	logger.out = os.Stdout
	logger.sync = sync
	logger.time = time.Now()
	logger.request = make(chan request, DefaultRequestSize)
	logger.quit = make(chan bool)

	if file != "" {
		out, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend|0644)
		if err != nil {
			panic(err)
		}

		logger.fileName = file
		logger.flog = out
		logger.out = io.MultiWriter(out, os.Stdout)
	}

	if !sync {
		go logger.watcher()
	}

	return logger
}

// Logger :
type Logger struct {
	out     io.Writer
	sync    bool
	time    time.Time
	request chan request
	quit    chan bool

	wlock    sync.Mutex // writer lock
	fileName string
	flog     *os.File
}

func (l *Logger) log(level string, name string, v ...interface{}) {
	if l.sync {
		t := time.Now()

		result := fmt.Sprintf(DefaultFormat, t.Format(DefaultTimeFormat), level, fmt.Sprint(v...))

		if l.flog != nil {
			d := 24 * time.Hour

			if t.Truncate(d).After(l.time.Truncate(d)) {
				l.rotate()
			}
		}

		l.wlock.Lock()
		defer l.wlock.Unlock()

		fmt.Fprintln(l.out, result)
	} else {
		r := new(request)
		r.time = time.Now()
		r.level = level
		r.name = name
		r.v = v

		l.request <- *r
	}
}

func (l *Logger) logf(level string, name string, format string, v ...interface{}) {
	if l.sync {
		t := time.Now()

		result := fmt.Sprintf(DefaultFormat, t.Format(DefaultTimeFormat), level, fmt.Sprintf(format, v...))

		if l.flog != nil {
			d := 24 * time.Hour

			if t.Truncate(d).After(l.time.Truncate(d)) {
				l.rotate()
			}
		}

		l.wlock.Lock()
		defer l.wlock.Unlock()

		fmt.Fprintln(l.out, result)
	} else {
		r := new(request)
		r.time = time.Now()
		r.level = level
		r.name = name
		r.format = format
		r.v = v

		l.request <- *r
	}
}

// Info : print info level log
func (l *Logger) Info(name string, v ...interface{}) {
	l.log(INFO, name, v...)
}

// Infof : print info level log with format
func (l *Logger) Infof(name string, format string, v ...interface{}) {
	l.logf(INFO, name, format, v...)
}

// Error : print error level log
func (l *Logger) Error(name string, v ...interface{}) {
	l.log(ERROR, name, v...)
}

// Errorf : print error level log with format
func (l *Logger) Errorf(name string, format string, v ...interface{}) {
	l.logf(ERROR, name, format, v...)
}

// Warn : print warning level log
func (l *Logger) Warn(name string, v ...interface{}) {
	l.log(WARN, name, v...)
}

// Warnf : print warning level log with format
func (l *Logger) Warnf(name string, format string, v ...interface{}) {
	l.logf(WARN, name, format, v...)
}

// Debug : print debug level log
func (l *Logger) Debug(name string, v ...interface{}) {
	l.log(DEBUG, name, v...)
}

// Debugf : print debug level log with format
func (l *Logger) Debugf(name string, format string, v ...interface{}) {
	l.logf(DEBUG, name, format, v...)
}

// Fatal : print fatal level log and exit program
func (l *Logger) Fatal(name string, v ...interface{}) {
	l.log(FATAL, name, v...)
	l.Destroy()
	os.Exit(2)
}

// Destroy sends quit signal to watcher and releases all the resources
func (l *Logger) Destroy() {
	if !l.sync {
		// quit watcher
		l.quit <- true
		// wait for watcher quit
		<-l.quit
	}
	// clean up
	if l.flog != nil {
		l.flog.Close()
	}
}

func (l *Logger) flushReq(req request) {
	var msg string
	if req.format != "" {
		msg = fmt.Sprintf(req.format, req.v...)
	} else {
		msg = fmt.Sprint(req.v...)
	}

	result := fmt.Sprintf(DefaultFormat, req.time.Format(DefaultTimeFormat), req.level, msg)

	if l.flog != nil {
		d := 24 * time.Hour

		if req.time.Truncate(d).After(l.time.Truncate(d)) {
			l.rotate()
		}
	}

	fmt.Fprintln(l.out, result)
}

func (l *Logger) watcher() {
	for {
		timeout := time.After(time.Millisecond * 100)

		select {
		case req := <-l.request:
			l.flushReq(req)
		case <-timeout:
			// do nothing
		case <-l.quit:
			// If quit signal received, cleans the channel
			// and writes all of them to io.Writer.
			for {
				select {
				case req := <-l.request:
					l.flushReq(req)
				default:
					l.quit <- true
					return
				}
			}
		}
	}
}

func (l *Logger) rotate() (err error) {

	l.wlock.Lock()
	defer l.wlock.Unlock()

	// Close existing file if open
	if l.flog != nil {
		err = l.flog.Close()
		l.flog = nil
		if err != nil {
			return
		}
	}
	// Rename dest file if it already exists
	_, err = os.Stat(l.fileName)
	if err == nil {
		err = os.Rename(l.fileName, l.fileName+"."+l.time.Format("2006-01-02")+".log")
		if err != nil {
			return
		}
	}

	l.time = time.Now()

	// Create a file.
	l.flog, err = os.Create(l.fileName)
	if err == nil {
		l.out = io.MultiWriter(l.flog, os.Stdout)
	}

	return
}
