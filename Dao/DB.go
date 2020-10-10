package Dao

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"icxl/Entitys"
)

var DB = GetDB()

func GetDB() *gorm.DB {
	dsn := "sqlserver://sa:Bosshong2012@111.229.26.207?database=gorm"
	db, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Entitys.Account{})
	InitData(db)



	return db
}
func InitData(db *gorm.DB)  {
	var count int64=0
	db.Model(Entitys.Account{}).Count(&count)
	if count==0{
		studentaccount :=Entitys.Account{Name: "icxl",PassWord: "123456"}
		db.Create(&studentaccount)
	}
}
