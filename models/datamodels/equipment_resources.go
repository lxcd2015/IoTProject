package datamodels

type EquipmentResources struct {
	Id      int
	Name    string
	Sdk     string
	Info    string
	Details string `orm:"type(text)"`
	TotalId int
	TypeId  int
}
