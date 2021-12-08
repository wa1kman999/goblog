package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/global"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"go.uber.org/zap"
)

// GetUserInfo 查询单个用户
func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		global.GBLog.Error("用户id为空")
		vs.SendParamParseError(ctx)
		return
	}
	id, _ := strconv.Atoi(userId)
	userInfo, err := userService.NewAppFormService().GetUserInfo(id)
	if err != nil {
		global.GBLog.Error("查询用户失败：", zap.Error(err))
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOkData(ctx, userInfo)
}
