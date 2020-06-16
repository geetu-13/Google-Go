package main

import "fmt"


func main(){
	/*
	var x int
	x=100 
	fmt.Printf("%d",x)
	*/

	/*
	var x int =1;
	var y int
	var ip *int
	ip=&x
	y=*ip
	fmt.Printf("%d",y)
	*/

	/*
	ptr:=new (int)
	*ptr=3
	fmt.Printf("%d",*ptr)
	*/

	/*
	fmt.Println("Hi")
	x:="Joe"
	fmt.Println("HI "+x)
	fmt.Printf("Hi %s",x)
	*/

	/*
	fmt.Printf("Number of apples?")

	var i int
    fmt.Scanf("%d", &i)
	fmt.Printf("%d",i)
	*/

	/*
	arr:=[...]string{"a","b","c","d","e","f","g","h"}
	s1:=arr[1:3]
	s2:=arr[2:5]
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	*/

	/*
	a1:=[...]string{"a","b","c"}
	sl:=a1[0:1]
	fmt.Printf("%d %d",len(sl), cap(sl))
	*/

	/*
	arr:=[...]string{"a","b","c","d","e","f","g","h"}
	s1:=arr[1:3]
	fmt.Println(s1[0])
	arr[1]="Def"
	fmt.Println(s1[0])
	*/

	/*
	i:=1
	defer fmt.Println(i+1)
	i++
	fmt.Println("Hello")
	*/

	defer fmt.Println("Bye!")
	fmt.Println("HEllo")

}