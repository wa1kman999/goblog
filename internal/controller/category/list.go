package category

import (
	"github.com/gin-gonic/gin"
	categoryService "github.com/wa1kman999/goblog/internal/application/service/category"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type Param struct {
	Name     string `json:"Name"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

// GetCategoryList 查询分类列表
func GetCategoryList(ctx *gin.Context) {
	var r Param
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	categoryList, count, err := categoryService.NewCategoryService().GetCategoryList(r.Name, r.Page, r.PageSize)
	if err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v查询分类列表失败", ctx.Request.RequestURI)
		vs.SendBad(ctx, err)
		return
	}
	res := vs.NewResData(r.Page, r.PageSize, count, categoryList)
	vs.SendOkData(ctx, res)
}
