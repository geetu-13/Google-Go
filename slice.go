package main

import (
	"fmt"
	"strconv"
	"unicode"
	"sort"
)
func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func sorting(temp [] int) {
	sort.Ints(temp)
	fmt.Println(temp)
}

func main(){
	sli:= make([]int, 3)
	var i int 
	for {
		fmt.Printf("Enter number: ")
		var v string
		fmt.Scanln(&v)
		if isInt(v){
			k,_:=strconv.Atoi(v)
			if (i==2){//exception case 
				sli[0]=k
			} else if (i>=3){
				sli=append(sli, k)
			} else {
				sli[i]=k
			}	
		} else if v=="X"{
			break
		} else {
			continue
		}
		i+=1
		
		fmt.Printf("slice: ")
		
		sorting(sli)
		
	}
}