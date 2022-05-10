package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"strconv"
)

//DeleteUser 删除
func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		logger.Errorf(ctx.Request.Context(), errors.New("userId为空"), "%v绑定参数失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	id, _ := strconv.Atoi(userId)
	err := userService.NewUserService().DeleteUser(id)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v删除信息失败", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
