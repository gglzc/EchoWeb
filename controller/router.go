package controller

import (
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"github.com/gglzc/echoWeb/controller/api"
)

func init(){

}

func Server(){
	e:=echo.New()
	//中間件
	zaplogger,_:=zap.NewProduction()
	e.Use(echozap.ZapLogger(zaplogger))
	//User的資料申請登入
	user:=e.Group("user")
	user.POST("register",controller.RegisterAccount)
	user.POST("login")
	//文章的情況
	post:=e.Group("post")

	e.Logger.Fatal(e.Start(":8080"))
}