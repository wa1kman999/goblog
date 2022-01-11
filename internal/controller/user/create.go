package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/wa1kman999/goblog/internal/application/service/user"
	"github.com/wa1kman999/goblog/internal/http/vs"
	//"github.com/sirupsen/logrus"
	"github.com/wa1kman999/goblog/pkg/common/logger"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

func CreateUser(ctx *gin.Context) {
	var user model.User
	log := logger.WithContext(ctx.Request.Context())
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Errorf(err, "user param is invalid")
		vs.SendParamParseError(ctx)
		return
	}
	if err := userService.NewAppFormService().CreateUser(&user); err != nil {
		log.Errorf(err, "create user failed")
		vs.SendBad(ctx, err)
		return
	}
	vs.SendOK(ctx)
}
