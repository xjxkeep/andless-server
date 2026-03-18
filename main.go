package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := LoadConfig()

	ossClient, err := NewOSSClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize OSS client: %v", err)
	}

	store := NewCaptchaStore()
	handler := NewCaptchaHandler(store, ossClient)

	r := gin.Default()
	api := r.Group("/api/captcha")
	{
		api.POST("/generate", handler.Generate)
		api.POST("/verify", handler.Verify)
	}

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
