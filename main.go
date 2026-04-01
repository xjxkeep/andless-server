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
	captchaHandler := NewCaptchaHandler(store, ossClient)
	versionHandler := NewVersionHandler(ossClient)

	r := gin.Default()
	captchaAPI := r.Group("/api/captcha")
	{
		captchaAPI.POST("/generate", captchaHandler.Generate)
		captchaAPI.POST("/verify", captchaHandler.Verify)
	}
	versionAPI := r.Group("/api/version")
	{
		versionAPI.GET("/check", versionHandler.Check)
		versionAPI.GET("/download", versionHandler.Download)
	}

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
