package util
import jsoniter "github.com/json-iterator/go"

type json struct{
	request	 string `json:"request"`
	response string `json:"response"`
}

func Marshal(interface{}) []byte{
	// 使用ConfigCompatibleWithStandardLibrary完全兼容官方库
	var s interface{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	result, _ := json.Marshal(s)
	return  result
	// output: {"id":1,"age":27,"gender":1,"name":"yuchanns","location":{"Country":"China","Province":"Guangdong","City":"Shenzhen","District":"Nanshan"}}
  }


