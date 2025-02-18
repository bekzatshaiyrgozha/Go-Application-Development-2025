package main
import "fmt"

func quadratic(arr[] int ,ch1 chan int){
	for i:=0;i<5;i++{
		ch1 <- arr[i] * arr[i]
	}
}
 
func main() {

	arr := make([]int,5)
	ch1 := make(chan  int,10)
	for i:=0;i<5;i++{
		fmt.Scan(&arr[i])
	}
	arr2 := make([]int,5)
	for i:=-0;i<5;i++{
		fmt.Scan(&arr2[i])
	}

	go quadratic(arr,ch1)
	go quadratic(arr2,ch1)

	
	

	for i:=0;i<10;i++{
		fmt.Println(<-ch1)
	}

}