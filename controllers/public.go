package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/iamMarkchu/goldfish/controllers/requests"
	"github.com/iamMarkchu/goldfish/models"
	"github.com/iamMarkchu/goldfish/services"
	"github.com/iamMarkchu/goldfish/tools"
	"net/http"
)

// PublicController operations for Public
type PublicController struct {
	BaseController
	*services.UserService
}

func (c *PublicController) Prepare() {
	c.UserService = services.NewUserService(models.NewUser())
}

// URLMapping ...
func (c *PublicController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
}

// @Title Login Api
// @Description user login
// @Param userName string true "用户名"
// @Param password string true "用户密码"
// @router /login [post]
func (c *PublicController) Login() {
	var (
		rl      requests.LoginRequest
		isValid bool
		valid   = validation.Validation{}
		err     error
		user    *models.UserWithToken
	)
	if err = c.ParseForm(&rl); err != nil {
		c.JsonReturn("登录接口解析参数错误:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	if isValid, err = valid.Valid(&rl); err != nil {
		c.JsonReturn("登录接口验证参数错误:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	} else if !isValid {
		c.JsonReturn("登录接口验证参数失败", tools.GetErrorMap(valid.Errors), http.StatusExpectationFailed)
		return
	}

	if user, err = c.UserService.Login(rl); err != nil {
		c.JsonReturn("登录接口处理错误:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	c.JsonReturn("登录接口", user, http.StatusOK)
}

// @Title Register Api
// @Description user login
// @Param userName string true "用户名"
// @Param password string true "用户密码"
// @Param rePassword string true "用户密码"
// @Param mobile string true "用户手机号"
// @router /register [post]
func (c *PublicController) Register() {
	var (
		rr      requests.RegisterRequest
		isValid bool
		valid   = validation.Validation{}
		err     error
		user    *models.User
	)
	if err = c.ParseForm(&rr); err != nil {
		c.JsonReturn("注册接口解析参数错误:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	if isValid, err = valid.Valid(&rr); err != nil {
		c.JsonReturn("注册接口验证参数错误:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	} else if !isValid {
		c.JsonReturn("注册接口验证参数失败", tools.GetErrorMap(valid.Errors), http.StatusExpectationFailed)
		return
	}
	if user, err = c.UserService.Register(rr); err != nil {
		c.JsonReturn("注册接口处理错误:"+err.Error(), nil, http.StatusExpectationFailed)
	} else {
		c.JsonReturn("注册接口", user, http.StatusOK)
	}
}
