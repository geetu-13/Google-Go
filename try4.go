package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type person struct {
	FName string
	LName string
}

func main() {
	maxLength := 20
	var p person
	fmt.Printf("Enter File name:")
	var x string
	fmt.Scan(&x)
	content, err := ioutil.ReadFile(x)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	words := strings.Fields(text)
	for i := 1; i < len(words); i += 2 {
		p.FName = words[i-1]
		if len(p.FName) > maxLength {
			p.FName = p.FName[:maxLength]
		}
		p.LName = words[i]
		if len(p.LName) > maxLength {
			p.LName = p.LName[:maxLength]
		}
		fmt.Println((i+1)/2, ":First Name:", p.FName, "\t\tLast Name:", p.LName)
	}
}
