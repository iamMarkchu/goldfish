package models

import "time"

type Role struct {
	Id      int
	Name    string    `orm:"column(role_name);"`
	Title   string    `orm:"description(权限描述)"`
	Status  uint8     `orm:"default(1);description(状态字段, 1 => 正常, 0 => 删除, 2 => 禁用)"`
	Created time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	Updated time.Time `orm:"auto_now;type(datetime);description(更新时间)"`
}
