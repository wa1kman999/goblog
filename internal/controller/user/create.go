package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

type CreateUserReq struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

func CreateUser(ctx *gin.Context) {
	var r CreateUserReq
	if err := ctx.ShouldBindJSON(&r); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v,参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	if err := userService.NewUserService().CreateUser(model.User{
		Username: r.UserName,
		Password: r.PassWord,
	}); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "创建用户失败,%v", r.UserName, r.PassWord)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
