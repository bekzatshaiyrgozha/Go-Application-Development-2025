package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr1 = `[
	{"id": 17, "username": "kasym", "phone": 0},
	{"id": "17", "address": "none", "company": "KBTU"}
]`

func main() {
	data := []byte(jsonStr1)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id":       42,
		"username": "someuser",
	}
	var user2i interface{} = user2
	result, _ := json.Marshal(user2i)
	fmt.Printf("json string from map:\n %s\n", string(result))
}
