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
	Id          int
	UserName    string    `orm:"unique;size(50);description(用户名)"`
	Password    string    `orm:"description(密码)" json:"-"`
	Mobile      string    `orm:"unique;description(手机号)"`
	Avatar      string    `orm:"description(头像)"`
	Description string    `orm:"type(text);description(用户描述)"`
	Role        string    `orm:"size(20);default(normal_users);description(角色)"`
	Status      uint8     `orm:"default(1);description(状态字段, 1 => 正常, 0 => 删除, 2 => 禁用)"`
	Created     time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Updated     time.Time `orm:"auto_now;type(datetime);description(更新时间)"`
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
