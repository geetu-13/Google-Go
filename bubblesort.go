package main
import "fmt"

func swap(a*int, b*int){
    temp:=*a
    *a=*b
    *b=temp
}
func bubblesort(sli []int){
  var i,j int
  for i<len(sli){
    j=i+1
    for j<len(sli){
      if (sli[i]>sli[j]){
        swap(&sli[i],&sli[j])
      }
      j+=1
    }
    i+=1
  }
  fmt.Print("After sorting: ")
  fmt.Println(sli)
}
func main(){
  sli:=make ([]int, 10)
  var i int 
  for i < cap(sli){
    var x int 
    fmt.Scan(&x)
    sli[i]=x
    i+=1
  }
  fmt.Print("Before sorting: ")
  fmt.Println(sli)//initially 
  bubblesort(sli)//sorting

}