package article

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/wa1kman999/goblog/internal/application/service/article"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/article/model"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

func CreateArticle(ctx *gin.Context) {
	var r model.Article
	if err := ctx.ShouldBindJSON(&r); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v,参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	if err := articleService.NewArticleService().CreateArticle(r); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "创建用户失败,%v")
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
