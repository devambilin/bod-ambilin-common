package micro

import (
	"context"
	"encoding/json"
	"errors"
	mclient "github.com/devambilin/bod-ambilin-common/client"
	cgrpc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/kubernetes"
	sgrpc "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/log"
	"go-micro.dev/v4/web"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"time"
)

type contextKey string

var (
	svcClient     client.Client
	serviceName   string
	svcConf       = make(map[string]string)
	configKey     = contextKey("configKey")
	serviceKey    = contextKey("serviceName")
	reqTimeoutKey = contextKey("requestTimeout")
	maxMsgSize    = contextKey("maxMsgSize")
)

func GetServiceName() string {
	return serviceName
}

func GetConfig() map[string]string {
	return svcConf
}

func Client() client.Client {
	return svcClient
}

func NewService(opts ...micro.Option) micro.Service {
	isKube := false
	keys := []string{}

	var ticker *time.Ticker
	done := make(chan bool)

	var svc micro.Service
	reg := registry.DefaultRegistry
	options := []micro.Option{
		micro.Version("latest"),
		micro.RegisterTTL(time.Second * 30),
		micro.RegisterInterval(time.Second * 15),
		micro.Flags(
			&cli.StringFlag{
				Name:  "service_name",
				Value: "",
				Usage: "the service name",
			},
		),
		micro.AfterStart(func() error {
			ticker = time.NewTicker(2 * time.Hour)

			go func() {
				for {
					select {
					case <-done:
						return
					case <-ticker.C:
						debug.FreeOSMemory()
					}
				}
			}()
			return nil
		}),
		micro.BeforeStop(func() error {
			ticker.Stop()
			done <- true

			return nil
		}),
	}
	options = append(options, opts...)
	svc = micro.NewService(options...)
	ckeys := svc.Options().Context.Value(configKey)
	if ckeys != nil {
		keys = ckeys.([]string)
	}

	if m := svc.Options().Context.Value(serviceKey); m != nil {
		serviceName = m.(string)
	}

	rto := client.DefaultRequestTimeout
	if m := svc.Options().Context.Value(reqTimeoutKey); m != nil {
		rto = m.(time.Duration)
	}

	mxMsgSize := cgrpc.DefaultMaxSendMsgSize
	if m := svc.Options().Context.Value(maxMsgSize); m != nil {
		mxMsgSize = m.(int) * 1024 * 1024 // to MegaByte
	}
	svc.Init(
		micro.Server(sgrpc.NewServer(
			sgrpc.MaxMsgSize((1024 * 1024 * mxMsgSize)),
		)),
		micro.Client(cgrpc.NewClient(
			client.RequestTimeout(rto),
			client.PoolSize(500),
			cgrpc.MaxSendMsgSize((1024*1024*mxMsgSize)),
		)),
		micro.Registry(reg),
		micro.Action(func(c *cli.Context) error {
			if serviceName == "" {
				serviceName = c.String("service_name")
				if serviceName == "" {
					panic("flag service_name is required")
				}
			}

			if isKube {
				match, _ := regexp.MatchString("^[a-z0-9\\-]*$", serviceName)
				if !match {
					return errors.New("service name must consist of lower case alphanumeric characters or '-'")
				}
			}

			/* load service configuration */
			if isKube {
				svcConf = GetEnvConf(keys...)
			} else {
				svcConf = GetFileConf(serviceName+".json", keys...)
			}

			return nil
		}),
	)
	svc.Server().Init(
		server.Name(serviceName),
	)
	svcClient = svc.Client()
	mclient.Init(svcClient)
	return svc
}

func NewWebService(opts ...web.Option) web.Service {

	isKube := false
	keys := []string{}

	var svcWeb web.Service
	var svcReg registry.Registry

	// define service by condition
	_, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount")
	if err == nil {
		svcReg = kubernetes.NewRegistry()
		isKube = true
	} else {
		svcReg = registry.DefaultRegistry
	}

	options := []web.Option{
		web.Version("latest"),
		web.RegisterTTL(time.Second * 30),
		web.RegisterInterval(time.Second * 15),
		web.Flags(
			&cli.StringFlag{
				Name:  "service_name",
				Value: "",
				Usage: "the service name",
			},
		),
	}

	// append user options
	options = append(options, opts...)

	svcWeb = web.NewService(options...)
	ckeys := svcWeb.Options().Context.Value(configKey)
	if ckeys != nil {
		keys = ckeys.([]string)
	}

	if m := svcWeb.Options().Context.Value(serviceKey); m != nil {
		serviceName = m.(string)
	}

	rto := client.DefaultRequestTimeout
	if m := svcWeb.Options().Context.Value(reqTimeoutKey); m != nil {
		rto = m.(time.Duration)
	}

	// init service
	svcWeb.Init(
		web.Registry(svcReg),
		web.Action(func(c *cli.Context) {

			if serviceName == "" {
				serviceName = c.String("service_name")
				if serviceName == "" {
					panic("flag service_name is required")
				}
			}

			match, _ := regexp.MatchString("^[a-z0-9\\-]*$", serviceName)
			if !match {
				panic("service name must consist of lower case alphanumeric characters or '-'")
			}

			/* load service configuration */
			if isKube {
				svcConf = GetEnvConf(keys...)
			} else {
				svcConf = GetFileConf(serviceName+".json", keys...)
			}
		}),
	)

	svcClient = cgrpc.NewClient(
		client.RequestTimeout(rto),
		client.PoolSize(500),
		cgrpc.MaxSendMsgSize((1024 * 1024 * 5)),
	)
	mclient.Init(svcClient)

	return svcWeb
}

func Config(keys ...string) micro.Option {
	return func(o *micro.Options) {
		o.Context = context.WithValue(o.Context, configKey, keys)
	}
}

// WebConfig : set config to be loaded
func WebConfig(keys ...string) web.Option {
	return func(o *web.Options) {
		o.Context = context.WithValue(o.Context, configKey, keys)
	}
}

// ServiceName : set service name
func ServiceName(serviceName string) micro.Option {
	return func(o *micro.Options) {
		o.Context = context.WithValue(o.Context, serviceKey, serviceName)
	}
}

// WebServiceName : set service name
func WebServiceName(serviceName string) web.Option {
	return func(o *web.Options) {
		o.Context = context.WithValue(o.Context, serviceKey, serviceName)
	}
}

// GetFileConf : get config from JSON file
func GetFileConf(fileName string, keys ...string) map[string]string {

	execPath, _ := os.Executable()
	pathfile := filepath.Dir(execPath) + "/" + fileName

	log.Info("load config file: " + pathfile)

	file, err := os.Open(pathfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := make(map[string]string)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	ret := make(map[string]string)

	for _, key := range keys {
		ret[key] = cfg[key]
	}

	return ret
}

// GetEnvConf : get config from environment variable
func GetEnvConf(keys ...string) map[string]string {

	ret := make(map[string]string)

	for _, key := range keys {
		ret[key] = os.Getenv(key)
	}

	return ret
}

// RequestTimeout set client request timeout
func RequestTimeout(d time.Duration) micro.Option {
	return func(o *micro.Options) {
		o.Context = context.WithValue(o.Context, reqTimeoutKey, d)
	}
}

// WebRequestTimeout set client request timeout
func WebRequestTimeout(d time.Duration) web.Option {
	return func(o *web.Options) {
		o.Context = context.WithValue(o.Context, reqTimeoutKey, d)
	}
}

// WebMaxSendMsgSize set client max send msg size
func WebMaxSendMsgSize(d int) web.Option {
	return func(o *web.Options) {
		o.Context = context.WithValue(o.Context, maxMsgSize, d)
	}
}

// WebMaxRecvMsgSize set client max receive msg size
func WebMaxRecvMsgSize(d int) web.Option {
	return func(o *web.Options) {
		o.Context = context.WithValue(o.Context, maxMsgSize, d)
	}
}

// SvcMaxSendMsgSize set client max send msg size
func SvcMaxSendMsgSize(d int) micro.Option {
	return func(o *micro.Options) {
		o.Context = context.WithValue(o.Context, maxMsgSize, d)
	}
}

// SvcMaxRecvMsgSize set client max receive msg size
func SvcMaxRecvMsgSize(d int) micro.Option {
	return func(o *micro.Options) {
		o.Context = context.WithValue(o.Context, maxMsgSize, d)
	}
}
