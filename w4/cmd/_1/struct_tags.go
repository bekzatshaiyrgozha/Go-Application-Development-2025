package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	ID       int `json:"user_id,string"`
	Username string
	Address  string `json:",omitempty"`
	Comnpany string `json:"-"`
}

func main() {
	u := &User1{
		ID:       1,
		Username: "someuser",
		Address:  "test",
		Comnpany: "KBTU",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("json string: %s\n", string(result))
}
