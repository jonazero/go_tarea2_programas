package main

import "fmt"

func main(){
	var base float64
	var altura float64
	fmt.Println("Ingresa la base: ")
	fmt.Scan(&base)
	fmt.Println("Ingresa la altura: ")
	fmt.Scan(&altura)
	fmt.Println("Area = ",(base*altura)/2)
}