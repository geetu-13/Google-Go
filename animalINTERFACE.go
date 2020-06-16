package main
import "fmt"
type Animal interface{
	Eat()
	Move() 
	Speak()
}
type bird struct{
	food, locomotion, noise string
}
type cow struct{
	food, locomotion, noise string
}
type snake struct{
	food, locomotion, noise string
}
func (b bird) Eat(){
	fmt.Println(b.food)
}
func (b bird) Move(){
	fmt.Println(b.locomotion)
}
func (b bird) Speak(){
	fmt.Println(b.noise)
}
func (c cow) Eat(){
	fmt.Println(c.food)
}
func (c cow) Move(){
	fmt.Println(c.locomotion)
}
func (c cow) Speak(){
	fmt.Println(c.noise)
}
func (s snake) Eat(){
	fmt.Println(s.food)
}
func (s snake) Move(){
	fmt.Println(s.locomotion)
}
func (s snake) Speak(){
	fmt.Println(s.noise)
}
func main(){
m:=make(map[string]Animal)
for{
	var task,name,typename string
	fmt.Print(">")
	fmt.Scanln(&task,&name,&typename)
	
	if(task=="newanimal"){
		switch typename{
		case "cow": 
			animal:=cow{food:"grass",locomotion:"walk",noise:"moo"}
			m[name]=animal
		case "snake":
			animal:=snake{food:"mice",locomotion:"slither",noise:"hsss"}
			m[name]=animal
		case "bird":
			animal:=bird{food:"worms",locomotion:"fly",noise:"peep"}
			m[name]=animal
		}
		fmt.Println("Created it!")
	}else if(task=="query"){
		animal:=m[name]
		switch typename{
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}else{
		continue;
	}
}
}