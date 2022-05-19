package service

import (
	"github.com/wa1kman999/goblog/pkg/user/dao"
	"github.com/wa1kman999/goblog/pkg/user/model"
)

type DomainUser interface {
	// Create 创建一个用户
	Create(user model.User) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.User, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.User, page, pageSize int64) ([]*model.User, int64, error)
	// Update 更新
	Update(value map[string]interface{}, query interface{}, args ...interface{}) error
	// Delete 删除
	Delete(query interface{}, args ...interface{}) error
}

// DomainUserService 用户领域服务
type DomainUserService struct {
}

func NewDomainUserService() DomainUser {
	return new(DomainUserService)
}

// Create 新建一个人
func (domain *DomainUserService) Create(user model.User) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Create(user)
}

// FindOne 查询一个
func (domain *DomainUserService) FindOne(fields string, query interface{}, args ...interface{}) (model.User, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return model.User{}, err
	}
	return entity.FindOne(fields, query, args)
}

// FindManyByPage 分页查询
func (domain *DomainUserService) FindManyByPage(fields string, query *model.User, page, pageSize int64) ([]*model.User, int64, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return nil, 0, err
	}
	return entity.FindManyByPage(fields, query, page, pageSize)
}

// Update 更新
func (domain *DomainUserService) Update(value map[string]interface{}, query interface{}, args ...interface{}) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Update(value, query, args)
}

// Delete 删除
func (domain *DomainUserService) Delete(query interface{}, args ...interface{}) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Delete(query, args)
}
