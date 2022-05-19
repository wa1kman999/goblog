package user

import (
	"github.com/go-errors/errors"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"github.com/wa1kman999/goblog/pkg/user/service"
)

type IUserService interface {
	Login(param *model.User) (model.User, error)
	CreateUser(param model.User) error
	GetUserInfo(id int) (model.User, error)
	GetUserList(userName string, role int, page, pageSize int64) ([]*model.User, int64, error)
	EditUser(param model.User) error
	DeleteUser(id int) error
}

type ServiceUser struct{}

func NewUserService() IUserService {
	return &ServiceUser{}
}

// Login 登陆
func (app *ServiceUser) Login(param *model.User) (model.User, error) {
	userService := service.NewDomainUserService()
	user, err := userService.FindOne("*",
		map[string]interface{}{
			"username": param.Username,
			"password": param.Password,
		})
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CreateUser 创建user
func (app *ServiceUser) CreateUser(param model.User) error {
	userService := service.NewDomainUserService()
	// 查询是否有同名的
	user, _ := userService.FindOne("id", "username = ?", param.Username)
	// 大于0 表示有同名的
	if user.ID > 0 {
		return errors.New("名字重复")
	}
	if err := userService.Create(param); err != nil {
		return err
	}
	return nil
}

// GetUserInfo 获取用户详细信息
func (app *ServiceUser) GetUserInfo(id int) (model.User, error) {
	userService := service.NewDomainUserService()
	userInfo, err := userService.FindOne("id, username, role", "id = ?", id)
	if err != nil {
		return model.User{}, err
	}
	return userInfo, nil
}

// GetUserList 获取user列表
func (app *ServiceUser) GetUserList(userName string, role int, page, pageSize int64) ([]*model.User, int64, error) {
	userService := service.NewDomainUserService()
	query := &model.User{
		Username: userName,
		Role:     role,
	}
	userList, count, err := userService.FindManyByPage("id, username, role", query, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return userList, count, nil
}

// EditUser 更新user信息
func (app *ServiceUser) EditUser(param model.User) error {
	userService := service.NewDomainUserService()
	return userService.Update(map[string]interface{}{
		"username": param.Username,
		"role":     param.Role,
	}, "id = ?", param.ID)
}

func (app *ServiceUser) DeleteUser(id int) error {
	return service.NewDomainUserService().Delete("id = ?", id)
}
