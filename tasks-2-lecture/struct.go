package main

import "fmt"

type Person struct{
	name string
	age int
}

func ExStuct(){
	p := Person{name:"Bekzat", age:20}

	fmt.Println(p.age,p.name)

}