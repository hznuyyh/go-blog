package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	UserName   string
	Password   string
	NickName   string
	CreateTime time.Time
	UpdateTime time.Time
}

func init()  {
	orm.RegisterModel(new(User))
}