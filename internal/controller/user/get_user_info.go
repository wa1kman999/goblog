package user

import (
	"errors"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/controller/user/resp"
	"github.com/wa1kman999/goblog/internal/http/vs"
)

// GetUserInfo 查询单个用户
func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		logger.Errorf(ctx.Request.Context(), errors.New("userId为空"), "%v绑定参数失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	id, _ := strconv.Atoi(userId)
	user, err := userService.NewUserService().GetUserInfo(id)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "查询用户失败: %v", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	userInfo := resp.UserResp{
		Id:       user.ID,
		UserName: user.Username,
		Role:     user.Role,
	}
	vs.SendOkData(ctx, userInfo)
}
