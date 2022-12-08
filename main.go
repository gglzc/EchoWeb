package main

import (
	"github.com/gglzc/echoWeb/controller"
	
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