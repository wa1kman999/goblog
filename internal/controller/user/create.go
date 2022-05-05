package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	"github.com/wa1kman999/goblog/pkg/user/model"

	//"github.com/sirupsen/logrus"
	"github.com/wa1kman999/goblog/pkg/common/logger"
)

type CreateUserReq struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

func CreateUser(ctx *gin.Context) {
	var r CreateUserReq
	log := logger.WithContext(ctx.Request.Context())
	if err := ctx.ShouldBindJSON(&r); err != nil {
		log.Errorf(err, "user param is invalid")
		vs.SendParamParseError(ctx)
		return
	}
	if err := userService.NewAppFormService().CreateUser(&model.User{
		Username: r.UserName,
		Password: r.PassWord,
	}); err != nil {
		log.Errorf(err, "create user failed")
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
