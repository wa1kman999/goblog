package user

import (
	"github.com/go-errors/errors"
	"github.com/wa1kman999/goblog/pkg/user/model"
	"github.com/wa1kman999/goblog/pkg/user/service"
)

type AppService struct{}

func NewAppFormService() *AppService {
	return &AppService{}
}

// Login 登陆
func (app *AppService) Login(param *model.User) (model.User, error) {
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
func (app *AppService) CreateUser(param *model.User) error {
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
func (app *AppService) GetUserInfo(id int) (model.User, error) {
	userService := service.NewDomainUserService()
	userInfo, err := userService.FindOne("id, username, role", "id = ?", id)
	if err != nil {
		return model.User{}, err
	}
	return userInfo, nil
}

// GetUserList 获取user列表
func (app *AppService) GetUserList(userName string, role int, pageIndex, pageSize int64) ([]*model.User, int64, error) {
	userService := service.NewDomainUserService()
	query := &model.User{
		Username: userName,
		Role:     role,
	}
	userList, count, err := userService.FindManyByPage("id, username, role", query, pageIndex, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return userList, count, nil
}

func (app *AppService) EditUser() {

}
