package client

import c "go-micro.dev/v4/client"

var (
	client c.Client
)

func Init(c c.Client) {
	client = c
}

func Get() c.Client {
	return client
}
