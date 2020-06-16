package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println("eat", a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println("makes", a.noise)
}

func prompt(sign string) []string {
	fmt.Print(sign, " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	txt := strings.Split(scanner.Text(), " ")
	return txt
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  Animal{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"},
	}
	for {
		input := prompt(">")
		animal := input[0]
		verb := strings.Title(input[1])
		_, found := animals[animal]
		if !found {
			fmt.Println("No data on", animal)
			continue
		}

		what := animals[animal]

		meth := reflect.ValueOf(&what).MethodByName(verb)
		if meth.IsValid() {
			fmt.Print("A ", animal, " ")
			meth.Call([]reflect.Value{})
		} else {
			fmt.Println("No method", verb, "on", animal)
		}
	}
}