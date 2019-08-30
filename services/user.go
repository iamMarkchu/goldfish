package services

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/iamMarkchu/goldfish/controllers/requests"
	"github.com/iamMarkchu/goldfish/models"
	"github.com/iamMarkchu/goldfish/tools"
	"github.com/iamMarkchu/goldfish/tools/token"
	"strconv"
)

type UserService struct {
	*models.User
}

func NewUserService(user *models.User) *UserService {
	return &UserService{User: user}
}

func (s *UserService) Login(rl requests.LoginRequest) (*models.UserWithToken, error)  {
	var err error

	s.User.UserName = rl.UserName
	// 查询用户名
	if err = s.User.GetUserByName(); err != nil {
		return nil, err
	} else if err == orm.ErrNoRows {
		return nil, errors.New("用户名不存在")
	}

	// 验证密码
	if s.User.Password != tools.MD5(rl.Password) {
		return nil, errors.New("密码错误")
	}

	// 验证状态
	if s.User.Status != models.StatusNormal {
		return nil, errors.New("该账号已被禁用或未激活")
	}
	auth := token.GetToken(strconv.Itoa(s.User.Id), s.User.Role)
	return models.NewUserWithToken(s.User, auth.Token, auth.ExpireIn), nil
}

func (s *UserService) Register(rr requests.RegisterRequest) (*models.User, error) {
	var err error
	if rr.Password != rr.RePassword {
		return nil, errors.New("两次密码不一致")
	}

	s.User.UserName = rr.UserName
	s.User.Mobile = rr.Mobile

	if err = s.User.GetUserByName(); err == nil {
		return nil, errors.New("用户名已注册")
	}

	if err = s.User.GetUserByMobile(); err == nil {
		return nil, errors.New("手机号已注册")
	}

	s.User.Password = tools.MD5(rr.Password)
	s.User.Status = models.StatusNormal
	s.User.Role = models.RoleNormal
	if err = s.User.Create(); err != nil {
		return nil, err
	}
	return s.User, nil
}

func (s *UserService) GetUserById(i int) (*models.User, error) {
	var err error
	s.User.Id = i
	if err = s.User.GetUserById(); err != nil {
		return nil, err
	}
	return s.User, nil
}
