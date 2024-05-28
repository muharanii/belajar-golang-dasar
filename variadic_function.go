package main

import "fmt"

func sumAll(numbers ...int)int{
	total := 0

	for _, numbers := range numbers{
		total += numbers
	}
	return total 
}

func main(){
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10,10,10,10))
	fmt.Println(sumAll(10,10,10,10,10))
}