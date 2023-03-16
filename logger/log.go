package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logWriter struct{}

func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05.000") + " [ INFO] - " + string(bytes))
}

type errWriter struct {
	f *os.File
}

func (w errWriter) Write(bytes []byte) (int, error) {

	errStr := time.Now().Format("2006-01-02 15:04:05.000") + " [ERROR] - " + string(bytes)

	fmt.Print(errStr)

	return w.f.WriteString(errStr)
}

// Global Logger
var (
	mlog  *Logger
	mlogS *Logger

	format     = "%-23s [%5s] - %s\n time,levelname,message"
	timeFormat = "2006-01-02 15:04:05.000"
)

func init() {
	mlog = SimpleLogger()
	mlogS = CustomLogger("", true)

	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

// ToFile .
func ToFile(filename string) {
	mlog = CustomLogger(filename, false)
	mlogS = CustomLogger(filename, true)
}

// FatalToFile .
func FatalToFile(errlog error, fname string) {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error write log to file: " + err.Error())
	}
	defer f.Close()

	logF := log.New(errWriter{f}, "", 0)
	logF.Fatal(errlog)
}

// Info .
func Info(msg string) {
	mlog.Info("log", msg)
}

// InfoS synchronously log
func InfoS(msg string) {
	mlogS.Info("log", msg)
}

// Infof formated log
func Infof(format string, v ...interface{}) {
	mlog.Infof("log", format, v...)
}

func Env(value map[string]string) {
	for k, v := range value {
		mlog.Infof("log", "%v: %v", k, v)
	}
}

// Warn warning log
func Warn(msg string) {
	mlog.Warn("log", msg)
}

// WarnS synchronously log
func WarnS(msg string) {
	mlogS.Warn("log", msg)
}

// Warnf formated log
func Warnf(format string, v ...interface{}) {
	mlog.Warnf("log", format, v...)
}

// Error .
func Error(msg string) {
	mlog.Error("log", msg)
}

// ErrorS synchronously log
func ErrorS(msg string) {
	mlogS.Error("log", msg)
}

// Errorf formated log
func Errorf(format string, v ...interface{}) {
	mlog.Errorf("log", format, v...)
}

// Fatal .
func Fatal(err error) {
	log.Fatal(err)
}

// Panic .
func Panic(msg string) {
	log.Panic(msg)
}
