package main

import "fmt"

func StudentAndGrade(){
	var name string
	var grade int
	var count int
	students:=make(map[string]int)

	fmt.Println("Enter number of students!")
	fmt.Scan(&count)

	for i:=0;i<count;i++{
		fmt.Println("Enter name and grade!")
		fmt.Scan(&name,&grade)
		students[name] = grade
	}

	fmt.Println("Students and theyr grades:")

	for student  , grade := range students{
		fmt.Printf(" %s: %d\n", student  , grade)
	}
	main()

}

func books(){
	var name string
	var price float32
	var count int
	books:=make(map[string]float32)
	
	fmt.Println("Enter number of books:")
	fmt.Scan(&count)
	for i:=0;i<count;i++{
		fmt.Printf("Enter %d  of book name and price:",i+1)
		fmt.Scan(&name,&price)

		books[name] = price
	}

	for book,price := range books{
		fmt.Printf("%s: %.1f\n",book,price)
	}
	main()
}