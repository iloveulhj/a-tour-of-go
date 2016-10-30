package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	//World = "한글"

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
