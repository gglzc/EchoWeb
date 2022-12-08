package model

import (
	"time"

	model "github.com/gglzc/echoWeb/model/postgres"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db  *gorm.DB
var RDB *redis.Client
//初始化postgres
func InitDB(){
	  const(
		dsn = "host=0.0.0.0 user=admin password=test dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	  )
	  //開啟
	  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	  if err!=nil{
		return
	  }
	  //migrate 
	  db.AutoMigrate(&model.User{},&model.Comment{},&model.Post{},&model.UserStatistic{})
	  //設定最大idle數和同時連線
	  sqlDB,err:=db.DB()
	  if err!=nil{
			panic(err)
	  }
	  sqlDB.SetMaxIdleConns(10)
	  sqlDB.SetMaxOpenConns(100)
	  sqlDB.SetConnMaxIdleTime(time.Hour)
	  
	  defer sqlDB.Close()
}
  //初始化redis
func InitRedis(){
	rdb:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:" ",
		DB: 0 ,
	})
    RDB=rdb
}
