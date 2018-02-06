package models

import (
	"time"
)

//文章
type Articles struct {
	Id         int64
	Title      string    `orm:"size(50)"`                                //标题
	Content    string    `orm:"size(5000)"`                              //内容
	UserId     int64                                                     //创建用户id
	ShowFigure string    `orm:"size(150)"`                               //展示图
	Attachment string                                                    //附件
	Author     string    `orm:"size(20)"`                                //作者
	Views      int64     `orm:"index"`                                   //浏览次数
	ReplyTime  time.Time `orm:"index"`                                   //最后回复时间
	Reply      int64                                                     //评论条数
	Labels     []*Labels `orm:"rel(m2m);rel_table(blog_article_labels)"` //多对多
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime);index"`       //创建时间
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`                 //更新时间
	Users      *Users    `orm:"rel(fk)"`                                 //设置一对多
}
