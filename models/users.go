package models

import (
	"time"

	"blog/helpers"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id        int64
	Name      string      `orm:"size(20)"`                          //昵称
	Avatar    string      `orm:"size(150)"`                         //头像
	Email     string      `orm:"size(50);unique"`                   //邮箱
	Password  string      `orm:"size(32)"`                          //密码
	Status    int8        `orm:"size(1)"`                           //状态 1正常
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime);index"` //创建时间
	UpdatedAt time.Time   `orm:"auto_now;type(datetime)"`           //更新时间
	Articles  []*Articles `orm:"reverse(many)"`                     //设置用户与文章一对多关系
	Labels    []*Labels   `orm:"reverse(many)"`                     //设置用户与标签一对多关系
}

func GetUserById(id int64) (u Users) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	qs.Filter("id", id).One(&u, "Id", "Email", "Password", "Status", "Name", "Avatar")
	return u
}

func GetUserByEmail(email string) (u Users) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	qs.Filter("email", email).One(&u, "Id", "Email", "Password", "Status", "Name", "Avatar")
	return u
}

func CreateUser(data Users) (int64, error) {
	user := new(Users)
	user.Name = data.Name
	user.Email = data.Email
	user.Password = helpers.GetMd5(data.Password)
	user.Status = 1

	id, err := orm.NewOrm().Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
