package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/caarlos0/env"
	yaml "gopkg.in/yaml.v2"
)

// Client represents a client instance to connect MySQL.
type Client struct {
	*sql.DB
}

// Config represents configuration parameters to connect MySQL.
type Config struct {
	User     string `yaml:"user" env:"MYSQL_USER"`
	Password string `yaml:"password" env:"MYSQL_PASSWORD"`
	Host     string `yaml:"host" env:"MYSQL_HOST"`
	Database string `yaml:"database" env:"MYSQL_DATABASE"`
}

// NewClient creates a new client instance from config
func NewClient(config Config) (*Client, error) {
	log.Println(config)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.User, config.Password, config.Host, config.Database,
	))
	if err != nil {
		return nil, err
	}

	return &Client{DB: db}, nil
}

// NewFromYaml compose a new Client instance from a yaml file.
func NewFromYaml(path string) (*Client, error) {
	var config Config
	if err := load(path, &config); err != nil {
		return nil, err
	}

	return NewClient(config)
}

// NewFromEnv compose a new Client instance from environmental variables.
func NewFromEnv() (*Client, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return NewClient(config)
}

func load(path string, v interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, v)
}
