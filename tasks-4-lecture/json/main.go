package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:age`

}


func main(){
	jsonStr := `{"name": "Bekzat","age":20 }`

	var user User 
	json.Unmarshal([]byte(jsonStr),&user)

	fmt.Println(user.Name,user.Age)
}