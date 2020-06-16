package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var animals map[string]Animal

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")

}

func (c Cow) Move() {
	fmt.Println("walk")

}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}
func main() {
	animals := make(map[string]Animal)
	in := bufio.NewReader(os.Stdin)
	for {
		var command string
		fmt.Print("> ")
		command, _ = in.ReadString('\n')
		command = command[:len(command)-1]
		tokens := strings.Split(command, " ")

		switch tokens[0] {
		case "newanimal":
			switch tokens[2] {
			case "cow":
				if animals[tokens[1]] != nil {
					fmt.Println("The animal with such name exists! Try other name")
					continue
				}
				animals[tokens[1]] = Cow{tokens[1]}
				fmt.Println("Created it!")
			case "snake":
				if animals[tokens[1]] != nil {
					fmt.Println("The animal with such name exists! Try other name")
					continue
				}
				animals[tokens[1]] = Snake{tokens[1]}
				fmt.Println("Created it!")
			case "bird":
				if animals[tokens[1]] != nil {
					fmt.Println("The animal with such name exists! Try other name")
					continue
				}
				animals[tokens[1]] = Bird{tokens[1]}
				fmt.Println("Created it!")
			}
		case "query":
			switch tokens[2] {
			case "eat":
				animals[tokens[1]].Eat()
			case "move":
				animals[tokens[1]].Move()
			case "speak":
				animals[tokens[1]].Speak()
			}

		}
	}
}