package category

import (
	"github.com/gin-gonic/gin"
	categoryService "github.com/wa1kman999/goblog/internal/application/service/category"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/category/model"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var r CreateCategoryReq
	if err := ctx.ShouldBindJSON(&r); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v,参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	if err := categoryService.NewCategoryService().CreateCategory(model.Category{
		Name: r.Name,
	}); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "创建分类失败,%v", r.Name)
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
