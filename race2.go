package main
import (
	"fmt"
	"time"
)
var x int =15
func f1(){	
	x=x+1000
	
	fmt.Println("f1",x)
}
func f2(){

	fmt.Println("f2",x)
}

func main(){
	go f1()
	go f2()
	

	time.Sleep(time.Second)
	fmt.Println("done")
}