package service

import (
	"github.com/wa1kman999/goblog/pkg/user/dao"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

type DomainUser interface {
	Create(user *model.User) error
}

// DomainUserService 用户领域服务
type DomainUserService struct {
}

func NewDomainUserService() DomainUser {
	return new(DomainUserService)
}

// Create 新建一个人
func (domain *DomainUserService) Create(user *model.User) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Create(user)
}
