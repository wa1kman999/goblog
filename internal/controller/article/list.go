package article

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/wa1kman999/goblog/internal/application/service/article"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type Param struct {
	Title    string `json:"title"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

// GetArticleList 查询文章列表
func GetArticleList(ctx *gin.Context) {
	var r Param
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	articleList, count, err := articleService.NewArticleService().GetArticleList(r.Title, r.Page, r.PageSize)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v查询文章列表失败", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	res := vs.NewResData(r.Page, r.PageSize, count, articleList)
	vs.SendOkData(ctx, res)
}
