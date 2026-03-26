package main

import "testing"

func TestInferOSSUseCname(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		endpoint string
		want     bool
	}{
		{
			name:     "custom domain with scheme",
			endpoint: "https://oss.andless.tech",
			want:     true,
		},
		{
			name:     "custom domain without scheme",
			endpoint: "oss.andless.tech",
			want:     true,
		},
		{
			name:     "aliyun public endpoint",
			endpoint: "https://oss-cn-shanghai.aliyuncs.com",
			want:     false,
		},
		{
			name:     "aliyun internal endpoint",
			endpoint: "oss-cn-shanghai-internal.aliyuncs.com",
			want:     false,
		},
		{
			name:     "empty endpoint",
			endpoint: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := inferOSSUseCname(tt.endpoint); got != tt.want {
				t.Fatalf("inferOSSUseCname(%q) = %v, want %v", tt.endpoint, got, tt.want)
			}
		})
	}
}
