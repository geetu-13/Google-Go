package main

import (
	"fmt"
	"math"
)

func main() {
	var acc, vel, dis, t float64
	fmt.Printf("Initial acceleration: ")
	fmt.Scan(&acc)

	fmt.Printf("Initial velocity: ")
	fmt.Scan(&vel)

	fmt.Printf("Initial displacement: ")
	fmt.Scan(&dis)

	fnDis := GenDisplaceFn(acc, vel, dis)

	fmt.Printf("Time for displacement: ")
	fmt.Scan(&t)
	fmt.Printf("Displacement travelled after %g seconds is %g meters", t, fnDis(t))
}

func GenDisplaceFn(i_acc, i_vel, i_dis float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return (0.5 * i_acc * math.Pow(time, 2)) + (i_vel * time) + i_dis
	}
	return fn
}