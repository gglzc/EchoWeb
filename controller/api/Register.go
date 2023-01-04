package controller

import (
	"github.com/gglzc/echoWeb/model"
	"github.com/gglzc/echoWeb/request"
	"github.com/labstack/echo/v4"
)
/*先查找快取看有沒有申請過帳號
若有則申請成功
若無則先查詢資料庫
還是無則存入資料庫中及更新快取
*/
func RegisterController(c echo.Context)error{
	var request request.RegisterRequest

	err:=c.Bind(&request)
	if err!=nil{
		return err
	}
	//check cache if user exist or not
	
	//Create a new user
	newUser:=model.User{
		Username: request.UserName,
		Password: request.Password,
		Email: request.Email,
		
		//status==false 代表還沒開通帳號
		Status: false,
	}
	return nil
}
	
