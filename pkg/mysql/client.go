package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/caarlos0/env"
	yaml "gopkg.in/yaml.v2"
)

// Config represents configuration parameters to connect MySQL.
type Config struct {
	User     string `yaml:"user" env:"MYSQL_USER"`
	Password string `yaml:"password" env:"MYSQL_PASSWORD"`
	Host     string `yaml:"host" env:"MYSQL_HOST"`
	Database string `yaml:"database" env:"MYSQL_DATABASE"`
}

func open(c Config) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		c.User, c.Password, c.Host, c.Database,
	))
}

// OpenFromYaml compose a new DB handler from a yaml file.
func OpenFromYaml(path string) (*sql.DB, error) {
	var config Config
	if err := load(path, &config); err != nil {
		return nil, err
	}

	return open(config)
}

// OpenFromEnv compose a new DB handler from environmental variables.
func OpenFromEnv() (*sql.DB, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return open(config)
}

func load(path string, v interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, v)
}
