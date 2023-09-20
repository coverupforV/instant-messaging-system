package service

import (
	"IM/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Summary index
// @Schemes
// @Description index
// @Tags index
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
