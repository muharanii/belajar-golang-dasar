package main

import "fmt"

func sayHelloTo(firsatName string, lastName string){
	fmt.Println("Hello", firsatName, lastName)
}

func main(){
	sayHelloTo("Ira", "Muharani")
}