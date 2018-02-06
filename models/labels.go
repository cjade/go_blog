package models

import (
	"time"
)

//标签
type Labels struct {
	Id            int
	Name          string      `orm:"size(20);unique"`                   //名称
	UserId        int64                                                 //创建用户id
	ArticlesCount int16                                                 //标签下一共有多少篇文章
	Articles      []*Articles `orm:"reverse(many)"`                     //反向
	CreatedAt     time.Time   `orm:"auto_now_add;type(datetime);index"` //创建时间
	UpdatedAt     time.Time   `orm:"auto_now;type(datetime)"`           //更新时间
	Users         *Users      `orm:"rel(fk)"`                           //设置一对多
}
