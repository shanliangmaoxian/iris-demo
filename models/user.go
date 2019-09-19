package models

type User struct {
	Id   int    `xorm:"not null pk autoincr comment('id') INT(11)" form:"id"`
	Name string `xorm:"comment('账号') VARCHAR(255)" form:"name"`
	Pwd  string `xorm:"comment('密码') VARCHAR(255)" form:"pwd"`
}
