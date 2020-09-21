package main

import "fmt"
func main(){
	var f float64
	fmt.Println("Ingresa los grados fahrenheit: ")
	fmt.Scan(&f)
	fmt.Println("Celcius = ", (f-32) * 5/9)
}