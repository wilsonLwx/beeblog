package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	UID             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册medel
	orm.RegisterModel(new(Category), new(Topic))
	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:mysql@tcp(127.0.0.1:3306)/beeblog?charset=utf8", 30,30)
}
