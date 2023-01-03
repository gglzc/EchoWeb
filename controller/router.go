package controller

import (
<<<<<<< HEAD
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"github.com/gglzc/echoWeb/controller/api"
=======
	"gglzc/EchoWeb/controller"

	"github.com/brpaz/echozap"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
>>>>>>> 5971af5 (improve whole structure)
)

func Server(){
	//instance
	e:=echo.New()
<<<<<<< HEAD
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
=======
	/*middleware*/
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//log
	zapLogger,_:=zap.NewProduction()
	e.Use(echozap.ZapLogger(zapLogger))
	//jwt

	//handler
	e.POST("/register",Register)
	e.POST("/login",Login)
>>>>>>> 5971af5 (improve whole structure)
}