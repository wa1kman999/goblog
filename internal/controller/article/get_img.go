package article

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

// GetImg 获取图片资源
func GetImg(ctx *gin.Context) {
	path := ctx.Param("path")
	if path == "" {
		logger.Errorf(ctx.Request.Context(), errors.New("路径为空"), "%v参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	ctx.File("fileDir/" + path)
}
