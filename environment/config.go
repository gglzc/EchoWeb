package environment

import(
	"github.com/spf13/viper"
)

type Config struct{
	Port	string
}

func init(){
	Instance:=viper.GetViper()
	Instance.SetConfigType("yaml")
	
}