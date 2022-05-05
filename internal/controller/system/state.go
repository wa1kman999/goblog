package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	systemService "github.com/wa1kman999/goblog/internal/application/service/system"
	"github.com/wa1kman999/goblog/internal/http/vs"
)

// GetServerInfo 查询服务器状态
func GetServerInfo(ctx *gin.Context) {
	logger := logrus.WithContext(ctx.Request.Context())
	serverInfo, err := systemService.NewSystemService().GetServerInfo()
	if err != nil {
		logger.Errorf("查询服务器状态失败: %s", err)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOkData(ctx, serverInfo)
}
