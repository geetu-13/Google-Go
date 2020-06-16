package main
import "fmt"
import "os"
import "bufio"
import "strings"

type name struct {
	fname string
	lname string
}
func main(){
	var filename string
	fmt.Print("enter filename: ")
	fmt.Scanln(&filename)
	f,_:=os.Open(filename)

	sli:=[]*name{}

	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		//fmt.Println(scanner.Text())

		n1:= new (name)
		n:=strings.Split(scanner.Text()," ")
		n1.fname,n1.lname=n[0],n[1]
		
		sli=append(sli, n1)

	}
	
	for i:=range(sli){
		n:=sli[i]
		fmt.Println(n.fname+" "+n.lname)
	}

}