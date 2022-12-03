package model
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )
  
func DBConn(){

	  dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	  if err!=nil{
		return
	  }
	  db.AutoMigrate()
}
  