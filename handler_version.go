package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gocache "github.com/patrickmn/go-cache"
)

type VersionInfo struct {
	Channel   string `json:"channel"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
}

type VersionHandler struct {
	ossClient *OSSClient
	cache     *gocache.Cache
}

func NewVersionHandler(ossClient *OSSClient) *VersionHandler {
	return &VersionHandler{
		ossClient: ossClient,
		cache:     gocache.New(5*time.Minute, 10*time.Minute),
	}
}

func (h *VersionHandler) Check(c *gin.Context) {
	channel := c.DefaultQuery("channel", "release")
	if channel != "release" && channel != "dev" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid channel, must be 'release' or 'dev'"})
		return
	}

	cacheKey := fmt.Sprintf("version:%s", channel)
	if cached, found := h.cache.Get(cacheKey); found {
		c.JSON(http.StatusOK, cached.(*VersionInfo))
		return
	}

	var ossPath string
	if channel == "release" {
		ossPath = "release/latest/version.json"
	} else {
		ossPath = "dev/latest/version.json"
	}

	data, err := h.ossClient.GetObject(ossPath)
	if err != nil {
		log.Printf("failed to get version.json from OSS (%s): %v", ossPath, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch version info"})
		return
	}

	var info VersionInfo
	if err := json.Unmarshal(data, &info); err != nil {
		log.Printf("failed to parse version.json (%s): %v", ossPath, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse version info"})
		return
	}

	h.cache.Set(cacheKey, &info, gocache.DefaultExpiration)
	c.JSON(http.StatusOK, info)
}
