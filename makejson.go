package main

import "fmt"
import "encoding/json"

func main(){
	var p map [string] string
	p=make(map[string]string)
	var a,b string
	fmt.Print("enter name: ")
	fmt.Scan(&a)
	fmt.Print("enter address: ")
	fmt.Scan(&b)
	p["name"]=a
	p["address"]=b
	fmt.Println(p)//displaying the map
	barr,_:=json.Marshal(p)
	fmt.Println(string(barr))//displaying the json object
}