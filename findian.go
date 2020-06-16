package main

import "fmt"
import "strings"

func main(){
	var s string 
	fmt.Scanf("%q",&s)
	//fmt.Scanf() with format %q  a double-quoted string safely escaped with Go syntax 
	//fmt.Println(s)
	str:= strings.ToLower(s)
	//fmt.Println(str)
	if strings.HasPrefix(str,"i") && strings.Contains(str,"a") && strings.HasSuffix(str,"n"){
		fmt.Println("Found")
	} else{
	fmt.Println("Not Found!")
	}
}