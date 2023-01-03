package main
<<<<<<< HEAD

import (
	"github.com/gglzc/echoWeb/controller"
=======
import ("github.com/swaggo/echo-swagger" )
func main(){
>>>>>>> 5971af5 (improve whole structure)
	
	"fmt"
	"github.com/spf13/viper"
)

func init(){
	Instance:=viper.GetViper()
	Instance.SetConfigName("config")
	Instance.SetConfigType("yml")
	Instance.AddConfigPath("./")

	err:=Instance.ReadInConfig()
	if err!=nil{
		panic(fmt.Errorf("Fatal error config file: %w \n",err))
	}
}
func main(){
	//Start Run Server
	controller.Server()

}