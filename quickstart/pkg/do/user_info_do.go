package do

type UserInfo struct {
	Id       int32  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Gender   int32  `gorm:"column:gender"`
	Passport string `gorm:"column:passport"`
	Address  string `gorm:"column:address"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
