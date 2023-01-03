package model

import (
	"github.com/go-redis/redis/v8"
	"context"

)

var RedisInstance *redis.Client

func init(){
	RedisInstance=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func CheckUserExits(username string){
	err:=RedisInstance.Set(ctx context.Context,"username",1,0)
}