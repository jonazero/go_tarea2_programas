package main

import "fmt"
import "math"
func main(){
	var radio float64
	fmt.Println("Ingresa el radio: ")
	fmt.Scan(&radio)
	fmt.Println("Area = ",math.Pi*math.Pow(radio,2))
}