package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/global"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"go.uber.org/zap"
)

func CreateUser(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		global.GBLog.Error("user param is invalid:", zap.Error(err))
		return
	}
	if err := userService.NewAppFormService().CreateUser(user); err != nil {
		global.GBLog.Error("create user failed:", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "创建用户成功",
	})
}
