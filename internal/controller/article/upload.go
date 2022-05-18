package article

import (
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/internal/application/service/article"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

func Upload(ctx *gin.Context) {
	header, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "接收文件失败,%v", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	path, err := article.NewArticleService().Upload(header)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "保存文件失败,%v", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOkData(ctx, struct {
		Path string `json:"path"`
	}{Path: path})
}
