package config

import "os"

type GrpcConfig struct {
	Url string
}

type Config struct {
	GrpcConfig
}

func (c *Config) readConfig() {
	grpcUrl := os.Getenv("GRPC_URL")
	//grpcUrl := "localhost:8080"
	c.GrpcConfig = GrpcConfig{Url: grpcUrl}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
