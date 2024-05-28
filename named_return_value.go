package main

import "fmt"

func getTriName() (firsatName, middleName, lastName string){
	// firsatName = "Ira"
	// middleName = "Rania"
	// lastName = "Muharani"

	// return firsatName, lastName, middleName

	return "Muharani", "Ira","Rania"
}

func main (){
	namaSatu, namaDua, namaTiga := getTriName()
	
	fmt.Println(namaSatu, namaDua, namaTiga)
}