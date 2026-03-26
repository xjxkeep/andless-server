package main

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	bucket *oss.Bucket
}

func NewOSSClient(cfg Config) (*OSSClient, error) {
	client, err := oss.New(
		cfg.OSSEndpoint,
		cfg.OSSAccessKeyID,
		cfg.OSSAccessKeySecret,
		oss.UseCname(cfg.OSSUseCname),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OSS client: %w", err)
	}
	bucket, err := client.Bucket(cfg.OSSBucket)
	if err != nil {
		return nil, fmt.Errorf("failed to get OSS bucket: %w", err)
	}
	return &OSSClient{bucket: bucket}, nil
}

// GetObject reads an object from OSS and returns its content.
func (o *OSSClient) GetObject(objectKey string) ([]byte, error) {
	body, err := o.bucket.GetObject(objectKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get OSS object %s: %w", objectKey, err)
	}
	defer body.Close()
	return io.ReadAll(body)
}

// SignURL generates a pre-signed download URL valid for 5 minutes.
func (o *OSSClient) SignURL(objectKey string) (string, error) {
	url, err := o.bucket.SignURL(objectKey, oss.HTTPGet, 300)
	if err != nil {
		return "", fmt.Errorf("failed to sign URL: %w", err)
	}
	return url, nil
}
