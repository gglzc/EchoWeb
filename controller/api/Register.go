package controller

import (
	"github.com/gglzc/echoWeb/model"
	"github.com/gglzc/echoWeb/request"
	"github.com/labstack/echo/v4"
)

/*如果沒有申請過帳號則寫入db*/

func RegisterAccount(c echo.Context)error{  
	//user的request
	var req request.RegisterRequest
	//如果存在帳號則不能申請
	if !CheckExist{
		return error.Error("這帳號已經存在了！")
	}
	//序列化資料
	c.Response(&req)
	user:=model.User
	 	 
}

//查有沒有存在帳號
func CheckExist(c echo.Context)bool{
	if 
}

func CheckEmail(){

}

func CheckUserName()  {
	
}

