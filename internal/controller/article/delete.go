package article

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/wa1kman999/goblog/internal/application/service/article"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"strconv"
)

func DeleteArticle(ctx *gin.Context) {
	param := ctx.Param("id")
	articleId, err := strconv.Atoi(param)
	if err != nil {
		vs.SendParamParseError(ctx)
	}
	if err := articleService.NewArticleService().DeleteArticle(articleId); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "删除文章失败,%v")
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
