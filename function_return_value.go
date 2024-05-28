package main

import "fmt"

func getHello(name string)string {
	hello := "Hello " + name
	return hello
}

func main (){
	result := getHello("Ira")
	fmt.Println(result)


	fmt.Println(getHello("Muharani"))
	fmt.Println("Ira")
}