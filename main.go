package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, Go!")
}
func main()  {
	distanceCount := 0
	click := 0
	sprintZone := 110
	for distanceCount < sprintZone {
		click++
		fmt.Println("click", click)
		distanceCount+=10
		fmt.Println("distanceCount", distanceCount)
	}


}
