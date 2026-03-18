package main

import "os"

type Config struct {
	OSSEndpoint        string
	OSSAccessKeyID     string
	OSSAccessKeySecret string
	OSSBucket          string
	ServerPort         string
}

func LoadConfig() Config {
	cfg := Config{
		OSSEndpoint:        os.Getenv("OSS_ENDPOINT"),
		OSSAccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
		OSSAccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
		OSSBucket:          os.Getenv("OSS_BUCKET"),
		ServerPort:         os.Getenv("SERVER_PORT"),
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}
	return cfg
}
