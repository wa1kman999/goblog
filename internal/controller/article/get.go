package article

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/wa1kman999/goblog/internal/application/service/article"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"strconv"
)

func GetArticle(ctx *gin.Context) {
	param := ctx.Param("id")
	articleId, err := strconv.Atoi(param)
	if err != nil {
		vs.SendParamParseError(ctx)
	}
	article, err := articleService.NewArticleService().GetArticle(articleId)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "获取单个文章失败,%v")
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOkData(ctx, article)
}
