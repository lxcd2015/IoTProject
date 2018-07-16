package startup

import (
	. "iotproject/models/datamodels"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/IoTProject")

	// register model
	orm.RegisterModel(new(TestModel))

	// create table
	orm.RunSyncdb("default", false, true)
}
