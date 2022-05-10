package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

//EditUser 编辑用户
func EditUser(ctx *gin.Context) {
	var r model.User
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	err = userService.NewUserService().EditUser(r)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v更新用户信息失败", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
