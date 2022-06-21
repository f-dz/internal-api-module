package data

import "github.com/api-abc/internal-api-module/rest"

type Client struct {
	*rest.Client
}

func New(rc *rest.Client) *Client {
	return &Client{rc}
}
