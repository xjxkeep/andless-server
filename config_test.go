package main

import "testing"

func TestLoadConfigReadsOSSCname(t *testing.T) {
	t.Setenv("OSS_ENDPOINT", "oss-cn-shanghai.aliyuncs.com")
	t.Setenv("OSS_ACCESS_KEY_ID", "id")
	t.Setenv("OSS_ACCESS_KEY_SECRET", "secret")
	t.Setenv("OSS_BUCKET", "bucket")
	t.Setenv("OSS_CNAME", "oss.andless.tech")

	cfg := LoadConfig()
	if cfg.OSSCname != "oss.andless.tech" {
		t.Fatalf("LoadConfig().OSSCname = %q, want %q", cfg.OSSCname, "oss.andless.tech")
	}
}
