package dal

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/src/config"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", config.DB.DbUser, config.DB.DbPwd, config.DB.DbUrls, config.DB.DbName)
	Log.Printf("call mysql dataSource:%v", dataSource)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	bizerror.Check(err)

	err = orm.RegisterDataBase("default", "mysql", dataSource, config.DB.DbMaxIdleConns, config.DB.DbMaxConns)
	bizerror.Check(err)

	orm.RegisterModelWithPrefix("t_goblog_",
		new(model.Articles),
		new(model.ArticlesDetails),
		new(model.Tag),
		new(model.ArticlesTag),
		new(model.Category),
		new(model.ArticlesCategory),
		new(model.Comment),
		new(model.User),
	)

	err = orm.RunSyncdb("default", config.DB.DbForce, true)
	bizerror.Check(err)
}