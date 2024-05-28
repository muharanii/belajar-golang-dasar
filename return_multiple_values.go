package main

import "fmt"

func getFulName() (string, string){
	return "Ira", "Muharani"
}

func main(){
	_, lastName := getFulName()

	fmt.Println( lastName)
}