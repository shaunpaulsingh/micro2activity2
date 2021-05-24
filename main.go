package main

import "math"

func main(){

}

func add(a int, b int) int{
	return a + b
}

func subtract(a int, b int) int{
	return a - b
}

func multiply(a int, b int) int{
	return a * b
}

func divide(a int, b int) float64{

	var result = float64(a) / float64(b)
	if result == math.Inf(10){
		result = 0.0
	}
	return result
}