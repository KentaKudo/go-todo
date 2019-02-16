package mysql

import "database/sql"

type Client struct {
	*sql.DB
}

type Config struct{}

func NewClient(config Config) *Client {
	return &Client{}
}
