package main
import "fmt"


func GenDisplaceFn(acc, iv, is float64) (func (float64)(float64)){
	f:=func(t float64)(float64){
		s:=is
		s+=iv*t
		x:=acc*t*t*0.5
		s+=x
		return s
	}
	return f;
}
func main(){
	var acc, iv, is, t float64

	fmt.Print("acceleration: ")
	fmt.Scanln(&acc)

	fmt.Print("initial velocity: ")
	fmt.Scanln(&iv)

	fmt.Print("initial displacement: ")
	fmt.Scanln(&is)

	fmt.Print("time: ")
	fmt.Scanln(&t)

	fn:=GenDisplaceFn(acc,iv,is)

	fmt.Printf("final displacement: ")
	fmt.Println(fn(t))



}