package main

import (
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	OSSEndpoint        string
	OSSAccessKeyID     string
	OSSAccessKeySecret string
	OSSBucket          string
	OSSUseCname        bool
	ServerPort         string
}

func LoadConfig() Config {
	ossEndpoint := os.Getenv("OSS_ENDPOINT")
	cfg := Config{
		OSSEndpoint:        ossEndpoint,
		OSSAccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
		OSSAccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
		OSSBucket:          os.Getenv("OSS_BUCKET"),
		OSSUseCname:        loadOSSUseCname(ossEndpoint),
		ServerPort:         os.Getenv("SERVER_PORT"),
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}
	return cfg
}

func loadOSSUseCname(endpoint string) bool {
	if raw := os.Getenv("OSS_USE_CNAME"); raw != "" {
		useCname, err := strconv.ParseBool(raw)
		if err == nil {
			return useCname
		}
	}

	return inferOSSUseCname(endpoint)
}

func inferOSSUseCname(endpoint string) bool {
	if endpoint == "" {
		return false
	}

	rawURL := endpoint
	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	host := strings.ToLower(parsed.Hostname())
	if host == "" {
		return false
	}

	return !strings.HasSuffix(host, ".aliyuncs.com") && host != "aliyuncs.com"
}
