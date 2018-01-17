package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:index`
	Views           int64     `orm:index`
	TopicTime       time.Time `orm:index`
	TopicCount      int64
	TopicLastUserId int64
}

//文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:size(5000)`
	Attachment      string
	Created         time.Time `orm:index`
	Update          time.Time `orm:index`
	Views           int64     `orm:index`
	Author          string
	ReplyTime       time.Time `orm:index`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("blog_db", "mysql", "root:123456@/blog_db?charset=utf8mb4", 10)
	orm.SetMaxOpenConns("blog_db", 30)
	orm.DefaultTimeLoc = time.UTC //默认时间
	orm.Debug = true
}
