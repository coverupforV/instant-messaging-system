package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex godoc
// @Summary 获取首页信息
// @Description 获取首页的详细信息
// @Tags index
// @Accept json
// @Produce json
// @Success 200 {string} Hello world
// @Router /index [get]

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
