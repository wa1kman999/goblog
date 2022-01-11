package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
)

type Param struct {
	UserName  string `json:"userName"`
	Role      int    `json:"role"`
	PageIndex int64  `json:"pageIndex"`
	PageSize  int64  `json:"pageSize"`
}

// GetUserList 查询用户列表
func GetUserList(ctx *gin.Context) {
	var param Param
	logger := logrus.WithContext(ctx.Request.Context())
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		logger.Error("查询用户列表参数保定失败")
		vs.SendParamParseError(ctx)
		return
	}
	userList, count, err := userService.NewAppFormService().GetUserList(param.UserName, param.Role, param.PageIndex, param.PageSize)
	if err != nil {
		logger.Errorf("查询用户列表失败: %s", err)
		vs.SendBad(ctx, err)
		return
	}
	res := vs.NewResData(param.PageIndex, param.PageSize, count, userList)
	vs.SendOkData(ctx, res)
}
