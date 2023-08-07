package category

import (
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type TestReq struct {
	Name string `json:"name" binding:"required"`
}

func Test(ctx *gin.Context) {
	var r TestReq
	if err := ctx.ShouldBindJSON(&r); err != nil {
		logger.Errorf(ctx.Request.Context(), err, "%v,参数绑定失败", ctx.Request.RequestURI)
		vs.SendParamParseError(ctx)
		return
	}
	vs.SendOkData(ctx, struct {
		Res string `json:"res"`
	}{Res: r.Name + "123123"})
}
