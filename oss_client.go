package main

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	bucket     *oss.Bucket
	signBucket *oss.Bucket
}

func NewOSSClient(cfg Config) (*OSSClient, error) {
	bucket, err := newBucket(cfg.OSSEndpoint, false, cfg)
	if err != nil {
		return nil, err
	}

	signBucket := bucket
	if cfg.OSSCname != "" {
		signBucket, err = newBucket(cfg.OSSCname, true, cfg)
		if err != nil {
			return nil, err
		}
	}

	return &OSSClient{
		bucket:     bucket,
		signBucket: signBucket,
	}, nil
}

func newBucket(endpoint string, useCname bool, cfg Config) (*oss.Bucket, error) {
	client, err := oss.New(
		endpoint,
		cfg.OSSAccessKeyID,
		cfg.OSSAccessKeySecret,
		oss.UseCname(useCname),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OSS client for endpoint %s: %w", endpoint, err)
	}

	bucket, err := client.Bucket(cfg.OSSBucket)
	if err != nil {
		return nil, fmt.Errorf("failed to get OSS bucket %s for endpoint %s: %w", cfg.OSSBucket, endpoint, err)
	}

	return bucket, nil
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
	url, err := o.signBucket.SignURL(objectKey, oss.HTTPGet, 300)
	if err != nil {
		return "", fmt.Errorf("failed to sign URL: %w", err)
	}
	return url, nil
}
