package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id           int        `json:"id"`
	Name         string    `orm:"description(类别名称)" json:"name"`
	AliasName    string    `orm:"description(类别别名)" json:"alias_name"`
	ParentCat    *Category `orm:"rel(fk);description(父类别);default(0);column(parent_id)" json:"parent_cat"`
	DisplayOrder uint32    `orm:"description(类别排序)" json:"display_order"`
	Image        string    `orm:"description(类别图片)" json:"image"`
	Status       uint8     `orm:"default(1);description(状态字段, 1 => 正常, 0 => 删除, 2 => 禁用)" json:"status"`
	User         *User     `orm:"rel(fk);description(创建人)" json:"user"`
	Created      time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"created"`
	Updated      time.Time `orm:"auto_now;type(datetime);description(更新时间)" json:"updated"`
}

func NewCategory() *Category {
	return &Category{}
}

func (m *Category) TableName() string {
	return "categories"
}

func (m *Category) Create() error {
	var (
		o = orm.NewOrm()
		err error
	)
	if _,err = o.Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Category) GetCategoryById() error {
	var (
		o = orm.NewOrm()
		err error
	)
	if err = o.Read(m); err == orm.ErrNoRows {
		return errors.New("查询不到")
	}
	return err
}

func (m *Category) Update() error {
	var (
		o = orm.NewOrm()
		err error
	)
	if err = o.Read(m); err == orm.ErrNoRows {
		return errors.New("查询不到")
	}
	return err
}
