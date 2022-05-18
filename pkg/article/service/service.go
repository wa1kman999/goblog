package service

import (
	"github.com/wa1kman999/goblog/pkg/article/dao"
	"github.com/wa1kman999/goblog/pkg/article/model"
)

type DomainArticle interface {
	// Create 创建一个用户
	Create(user model.Article) error
	// FindOne 查询一个
	FindOne(fields string, query interface{}, args ...interface{}) (model.Article, error)
	// FindManyByPage 分页查询
	FindManyByPage(fields string, query *model.Article, pageIndex, pageSize int64) ([]*model.Article, int64, error)
	// Update 更新
	Update(value map[string]interface{}, query interface{}, args ...interface{}) error
	// Delete 删除
	Delete(query interface{}, args ...interface{}) error
}

// DomainArticleService 文章
type DomainArticleService struct {
}

func NewDomainUserService() DomainArticle {
	return new(DomainArticleService)
}

// Create 新建一个人
func (domain *DomainArticleService) Create(user model.Article) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Create(user)
}

// FindOne 查询一个
func (domain *DomainArticleService) FindOne(fields string, query interface{}, args ...interface{}) (model.Article, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return model.Article{}, err
	}
	return entity.FindOne(fields, query, args)
}

// FindManyByPage 分页查询
func (domain *DomainArticleService) FindManyByPage(fields string, query *model.Article, pageIndex, pageSize int64) ([]*model.Article, int64, error) {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return nil, 0, err
	}
	return entity.FindManyByPage(fields, query, pageIndex, pageSize)
}

// Update 更新
func (domain *DomainArticleService) Update(value map[string]interface{}, query interface{}, args ...interface{}) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Update(value, query, args)
}

// Delete 删除
func (domain *DomainArticleService) Delete(query interface{}, args ...interface{}) error {
	entity, err := dao.NewUserEntity()
	if err != nil {
		return err
	}
	return entity.Delete(query, args)
}
