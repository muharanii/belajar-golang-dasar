package main

import "fmt"

func getGoodBye(name string)string{

	hello := "Hallo" + name
	return hello
}

func main() {
	goodBye := getGoodBye

	fmt.Println(goodBye("Muharani"))
}