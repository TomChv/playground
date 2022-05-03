package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	messenger "grpc-workshop/messenger"
)

type Client struct {
	conn *grpc.ClientConn
	core messenger.MessengerServiceClient
}

func New(conn *grpc.ClientConn) *Client {
	core := messenger.NewMessengerServiceClient(conn)
	conn.Connect()

	return &Client{
		conn, core,
	}
}

func (c *Client) Send(ctx context.Context, body string) (*messenger.Content, error) {
	return c.core.Send(ctx, &messenger.Content{Body: body, Date: uint64(time.Now().Unix())})
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}
