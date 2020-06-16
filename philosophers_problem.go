package main
import(
	"fmt"
	"sync"
	
)
type ChopS struct{sync.Mutex}
type Philo struct{
	leftCS, rightCS *ChopS
}
var ws sync.WaitGroup


func (p Philo) eat(x int ){
	for i:=0;i<3;i++{
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ",(x+1))
		fmt.Println("finished eating ",(x+1))

		p.leftCS.Unlock()
		p.rightCS.Unlock()

	}
	ws.Done()
}

func main(){
	CSticks:=make([]*ChopS,5)
	
	for i:=0; i<5; i++{
		CSticks[i]=new(ChopS)
	}
	philos:=make([]*Philo,5)
	for i:=0; i<5; i++{
		
		philos[i]=&Philo{CSticks[i],CSticks[(i+1)%5]}
		ws.Add(1)
		go philos[i].eat(i)
	}
	
	
	ws.Wait()

}