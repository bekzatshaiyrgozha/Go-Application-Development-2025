// go build gen/* && ./codegen.exe pack/packer.go  pack/marshaller.go
package main

import "fmt"

// lets generate code for this struct
// cgen: binpack
type User struct {
	ID       int
	RealName string `cgen:"-"`
	Login    string
	Flags    int
}

type Avatar struct {
	ID  int
	Url string
}

var test = 42

func main() {
	data := []byte{
		128, 36, 17, 0,

		4, 0, 0, 0,
		75, 66, 84, 85,

		16, 0, 0, 0,
	}

	u := User{}
	u.Unpack(data)
	fmt.Printf("Unpacked user %#v", u)
}
