package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/global"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"go.uber.org/zap"
)

func CreateUser(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		global.GBLog.Error("user param is invalid:", zap.Error(err))
		vs.SendParamParseError(ctx)
		return
	}
	if err := userService.NewAppFormService().CreateUser(user); err != nil {
		global.GBLog.Error("create user failed:", zap.Error(err))
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
