package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

const StatusNormal = 1
const RoleNormal = "normal_users"
const RoleAdmin = "admin"
const RoleOperator = "operators"

type User struct {
	Id          int        `json:"id"`
	UserName    string    `orm:"unique;size(50);description(用户名)" json:"user_name"`
	Password    string    `orm:"description(密码)" json:"-"`
	Mobile      string    `orm:"unique;description(手机号)" json:"mobile"`
	Avatar      string    `orm:"description(头像)" json:"avatar"`
	Description string    `orm:"type(text);description(用户描述)" json:"description"`
	Role        string    `orm:"size(20);default(normal_users);description(角色)" json:"role"`
	Status      uint8     `orm:"default(1);description(状态字段, 1 => 正常, 0 => 删除, 2 => 禁用)" json:"status"`
	Created     time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"created"`
	Updated     time.Time `orm:"auto_now;type(datetime);description(更新时间)" json:"updated"`
}

type UserWithToken struct {
	*User
	Token    string
	ExpireAt string
}

func NewUserWithToken(user *User, token string, expireAt string) *UserWithToken {
	return &UserWithToken{User: user, Token: token, ExpireAt: expireAt}
}

func NewUser() *User {
	return &User{}
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Create() error {
	var (
		o   = orm.NewOrm()
		err error
	)
	if _, err = o.Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) GetUserByName() error {
	var (
		o   = orm.NewOrm()
		err error
	)
	if err = o.Read(m, "UserName"); err == orm.ErrNoRows {
		return errors.New("查询不到")
	}
	return nil
}

func (m *User) GetUserByMobile() error {
	var (
		o   = orm.NewOrm()
		err error
	)
	if err = o.Read(m, "Mobile"); err == orm.ErrNoRows {
		return errors.New("查询不到")
	}
	return nil
}

func (m *User) GetUserById() error {
	var (
		o   = orm.NewOrm()
		err error
	)
	if err = o.Read(m); err == orm.ErrNoRows {
		return errors.New("查询不到")
	}
	return nil
}
