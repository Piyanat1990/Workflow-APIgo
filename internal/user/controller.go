package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Piyanat1990/workflow/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB, secret string) Controller {
	return Controller{
		Service: NewService(db, secret),
	}
}

func (controller Controller) Login(ctx *gin.Context) {
	var (
		request model.RequestLogin
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := controller.Service.Login(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"token",
		fmt.Sprintf("Bearer %v", token), int(10*time.Second),
		"/",
		"localhost",
		false,
		true,
)





	ctx.JSON(http.StatusOK, gin.H{
		"message": "login succeeed",
	})
}
