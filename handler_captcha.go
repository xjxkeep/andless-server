package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaHandler struct {
	captcha   *base64Captcha.Captcha
	ossClient *OSSClient
}

func NewCaptchaHandler(store *CaptchaStore, ossClient *OSSClient) *CaptchaHandler {
	driver := base64Captcha.NewDriverString(
		60,  // height
		240, // width
		0,   // noiseCount
		0,   // showLineOptions
		4,   // length
		"abcdefghijklmnopqrstuvwxyz0123456789",
		nil, // bgColor
		nil, // fonts
		nil, // fontsArray
	)
	captcha := base64Captcha.NewCaptcha(driver, store)
	return &CaptchaHandler{
		captcha:   captcha,
		ossClient: ossClient,
	}
}

type GenerateResponse struct {
	RequestID    string `json:"request_id"`
	CaptchaImage string `json:"captcha_image"`
}

type VerifyRequest struct {
	RequestID string `json:"request_id"`
	Result    string `json:"result"`
	Scene     string `json:"scene"`
}

type VerifyResponse struct {
	Success     bool   `json:"success"`
	DownloadURL string `json:"download_url,omitempty"`
	Message     string `json:"message,omitempty"`
}

func (h *CaptchaHandler) Generate(c *gin.Context) {
	id, b64s, answer, err := h.captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate captcha"})
		return
	}
	log.Println("answer", answer)
	c.JSON(http.StatusOK, GenerateResponse{
		RequestID:    id,
		CaptchaImage: b64s,
	})
}

func (h *CaptchaHandler) Verify(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, VerifyResponse{
			Success: false,
			Message: "请求参数错误",
		})
		return
	}

	if !h.captcha.Verify(req.RequestID, req.Result, true) {
		c.JSON(http.StatusOK, VerifyResponse{
			Success: false,
			Message: "验证码错误",
		})
		return
	}

	objectKey := GetObjectKey(req.Scene)
	if objectKey == "" {
		c.JSON(http.StatusOK, VerifyResponse{
			Success: false,
			Message: "无效的 scene 参数",
		})
		return
	}

	downloadURL, err := h.ossClient.SignURL(objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyResponse{
			Success: false,
			Message: "生成下载链接失败",
		})
		return
	}

	c.JSON(http.StatusOK, VerifyResponse{
		Success:     true,
		DownloadURL: downloadURL,
	})
}
