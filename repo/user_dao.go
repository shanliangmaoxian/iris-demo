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

func (d *UserDao) Get(id int) *models.User {
	data := &models.User{Id: id}
	if ok, err := d.engine.Get(data); ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserDao) Create(data *models.User) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *UserDao) Update(data *models.User) error {
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDao) Delete(id int) error {
	user := new(models.User)
	_, err := d.engine.Id(id).Delete(user)
	return err
}
