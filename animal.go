package main

import "fmt"


type Animal struct{
	food string
	locomotion string
	noise string
}
 
func (x Animal) Eat(){
	fmt.Println(x.food)
}
func (x Animal) Move(){
	fmt.Println(x.locomotion)
}
func (x Animal) Speak(){
	fmt.Println(x.noise)
}
func (x Animal) check (task string){
	if (task=="eat"){
		x.Eat()
	}else if (task=="move"){
		x.Move()
	}else if (task=="speak"){
		x.Speak()
	}
}


func main(){
	cow:=Animal{"grass","walk","moo"}
	bird:=Animal{"worms","fly","peep"}
	snake:=Animal{"mice","slither","hsss"}

	for{
		var n[2] string
		fmt.Print(">")
		fmt.Scanln(&n[0],&n[1])
		if (n[0]=="cow"){
			cow.check(n[1])
		}else if (n[0]=="bird"){
			bird.check(n[1])
		}else if(n[0]=="snake"){
			snake.check(n[1])
		}else{
			continue;
		}

	}


}