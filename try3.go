package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var name, address string
	m := make(map[string]string)

	fmt.Println("Please enter a name : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter an address : ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	address = input.Text()

	m[name] = address

	jsonStr, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonObject := string(jsonStr)

	fmt.Println("JSON object : ", jsonObject)
}
