package main

import "fmt"
import "math"
func main(){
	var lado float64
	fmt.Println("Ingresa el lado: ")
	fmt.Scan(&lado)
	fmt.Println("Area = ",math.Pow(lado,2))
}