package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/controller/user/resp"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type Param struct {
	UserName  string `json:"userName"`
	Role      int    `json:"role"`
	PageIndex int64  `json:"pageIndex"`
	PageSize  int64  `json:"pageSize"`
}

// GetUserList 查询用户列表
func GetUserList(ctx *gin.Context) {
	var r Param
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	userList, count, err := userService.NewUserService().GetUserList(r.UserName, r.Role, r.PageIndex, r.PageSize)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v查询用户列表失败", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	userRespList := make([]resp.UserResp, 0, len(userList))
	for _, v := range userList {
		temp := resp.UserResp{
			Id:       v.ID,
			UserName: v.Username,
			Role:     v.Role,
		}
		userRespList = append(userRespList, temp)
	}
	res := vs.NewResData(r.PageIndex, r.PageSize, count, userRespList)
	vs.SendOkData(ctx, res)
}
