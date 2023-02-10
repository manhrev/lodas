package config

import (
	"os"
	"strconv"
)

const (
	db_username = "DB_USERNAME"
	db_password = "DB_PWD"
	db_domain   = "DB_DOMAIN"
	db_port     = "DB_PORT"
	db_name     = "DB_NAME"

	grpc_host    = "GRPC_HOST"
	grpc_port    = "GRPC_PORT"
	grpc_network = "GRPC_NETWORK"

	cache_host     = "CACHE_HOST"
	cache_port     = "CACHE_PORT"
	cache_password = "CACHE_PWD"
	cache_db       = "CACHE_DB"
)

type deployConfig struct {
}

func (c *deployConfig) New() (*Configuration, error) {
	conf := &Configuration{}

	conf.Database = Database{
		Username: os.Getenv(db_username),
		Password: os.Getenv(db_password),
		Domain:   os.Getenv(db_domain),
		Port:     os.Getenv(db_port),
		Name:     os.Getenv(db_name),
	}
	cacheDb, err := strconv.Atoi(os.Getenv(cache_db))
	if err != nil {
		return nil, err
	}
	conf.Cache = Cache{
		Host:     os.Getenv(cache_host),
		Port:     os.Getenv(cache_port),
		Password: os.Getenv(cache_password),
		Db:       cacheDb,
	}

	conf.Grpc = Grpc{
		Host:    os.Getenv(grpc_host),
		Network: os.Getenv(grpc_network),
		Port:    os.Getenv(grpc_port),
	}

	return conf, nil
}
