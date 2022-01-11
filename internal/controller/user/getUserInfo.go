package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/controller/user/resp"
	"github.com/wa1kman999/goblog/internal/http/vs"
)

// GetUserInfo 查询单个用户
func GetUserInfo(ctx *gin.Context) {
	logger := logrus.WithContext(ctx.Request.Context())
	userId := ctx.Param("id")
	if userId == "" {
		logger.Error("用户id为空")
		vs.SendParamParseError(ctx)
		return
	}
	id, _ := strconv.Atoi(userId)
	user, err := userService.NewAppFormService().GetUserInfo(id)
	if err != nil {
		logger.Errorf("查询用户失败: %s", err)
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
