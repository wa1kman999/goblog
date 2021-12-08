package service

import (
	"github.com/wa1kman999/goblog/pkg/user/dao"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

type DomainUser interface {
	// Create 创建一个用户
	Create(user *model.User) error
	// FindOne 查询一个
	FindOne(fields []string, query interface{}) (*model.User, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.User, pageIndex, pageSize int64) ([]*model.User, int64, error)
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

// FindOne 查询一个
func (domain *DomainUserService) FindOne(fields []string, query interface{}) (*model.User, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return nil, err
	}
	return entity.FindOne(fields, query)
}

// FindManyByPage 分页查询
func (domain *DomainUserService) FindManyByPage(fields string, query *model.User, pageIndex, pageSize int64) ([]*model.User, int64, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return nil, 0, err
	}
	return entity.FindManyByPage(fields, query, pageIndex, pageSize)
}
