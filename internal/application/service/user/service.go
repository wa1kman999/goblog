package user

import (
	"github.com/wa1kman999/goblog/pkg/user/model"
	"github.com/wa1kman999/goblog/pkg/user/service"
)

type AppService struct{}

func NewAppFormService() *AppService {
	return &AppService{}
}

// CreateUser 创建user
func (app *AppService) CreateUser(param *model.User) error {
	userService := service.NewDomainUserService()
	if err := userService.Create(param); err != nil {
		return err
	}
	return nil
}
