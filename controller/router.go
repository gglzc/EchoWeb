package controller

import (
	"gglzc/EchoWeb/controller"

	"github.com/brpaz/echozap"
	controller "github.com/gglzc/echoWeb/controller/api"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

func Server(){
	//instance
	e:=echo.New()
	/*middleware*/
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//log
	zapLogger,_:=zap.NewProduction()
	e.Use(echozap.ZapLogger(zapLogger))
	//jwt

	//handler
	e.POST("/register",controller.RegisterController)
	e.POST("/login",Login)
}