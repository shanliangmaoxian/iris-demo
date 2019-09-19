package repo

import (
	"github.com/go-xorm/xorm"
	"iris-demo/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine: engine}
}

func (d *UserDao) GetAll() []models.User {
	datalist := make([]models.User, 0)
	if err := d.engine.Desc("id").Find(&datalist); err != nil {
		return datalist
	} else {
		return datalist
	}
}
