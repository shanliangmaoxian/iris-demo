package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/liukunxin/superstar-db/conf"
	"iris-demo/utils"
	"log"
	"sync"
)

var (
	masterEngine *xorm.Engine
	lock         sync.Mutex
)

//主库，单例
func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	config, _ := utils.GetYaml()
	c := config.Mysql
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.Dbname)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatalf("dbhelper.DbInstanceMaster,", err)
		return nil
	}

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	engine.SetTZDatabase(conf.SysTimeLocation)

	masterEngine = engine
	return engine
}
