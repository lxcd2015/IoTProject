package startup

import (
	. "iotproject/models/datamodels"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("ConnectionStrings"))

	// register model
	orm.RegisterModel(new(EquipmentResources))
	orm.RegisterModel(new(Functions))
	orm.RegisterModel(new(Total))
	orm.RegisterModel(new(Type))

	// create table
	orm.RunSyncdb("default", false, true)
}
