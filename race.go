package main
import "fmt"
import "time"
func f(s string){
	fmt.Println(s)
}
func main(){
	go f("first")
	go f("second")
	go f("third")


	time.Sleep(time.Second)
	fmt.Println("done")
}