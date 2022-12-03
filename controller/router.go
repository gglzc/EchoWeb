package controller

import(
	"github.com/labstack/echo/v4"
)

func init(){

}

func Server(){
	e:=echo.New()
	e.Use(e.Logger())
	e.POST("/register",Register)
	e.POST("/login",Login)
}