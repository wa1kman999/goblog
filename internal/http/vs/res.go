package vs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Res api response结构
type Res struct {
	Success bool        `json:"success"`
	Code    int         `json:"code,omitempty"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ResData 返回值data
type ResData struct {
	PageIndex int64       `json:"pageIndex"` // 当前页面
	PageSize  int64       `json:"pageSize"`  // 当前页大小
	Count     int64       `json:"count"`     // 总条数
	Data      interface{} `json:"data"`      // 数据
}

// NewRes 新建api response
func NewRes(success bool, msg string, data interface{}) *Res {
	res := Res{}
	res.Success = success
	res.Msg = msg
	res.Data = data
	return &res
}

// DefaultRes 获取默认api response
func DefaultRes() *Res {
	res := Res{}
	res.Success = true
	res.Msg = SUCCESS
	res.Data = struct{}{}
	return &res
}

// NewResData return new resData
func NewResData(PageIndex, pageSize, total int64, resData interface{}) ResData {
	return ResData{
		PageIndex: PageIndex,
		PageSize:  pageSize,
		Count:     total,
		Data:      resData,
	}
}

// SendOK 成功
func SendOK(ctx *gin.Context) {
	res := NewRes(true, SUCCESS, struct{}{})
	res.Code = StatusOK
	ctx.JSON(http.StatusOK, res)
}

// SendOkData 返回数据
func SendOkData(ctx *gin.Context, resData interface{}) {
	res := NewRes(true, SUCCESS, struct{}{})
	res.Code = StatusOK
	res.Data = resData
	ctx.JSON(http.StatusOK, res)
}

// SendParamParseError 返回参数解析失败错误
func SendParamParseError(ctx *gin.Context) {
	res := NewRes(false, ReqDataValError, struct{}{})
	res.Code = StatusBadRequest
	ctx.JSON(http.StatusBadRequest, res)
}

// SendBad 失败
func SendBad(ctx *gin.Context, err error) {
	res := NewRes(false, err.Error(), struct{}{})
	res.Code = StatusBadRequest
	ctx.JSON(http.StatusBadRequest, res)
}

// SendBadData 返回错误
func SendBadData(ctx *gin.Context, err error, resData interface{}) {
	res := NewRes(false, err.Error(), struct{}{})
	res.Code = StatusBadRequest
	res.Data = resData
	ctx.JSON(http.StatusBadRequest, res)
}
