package main

import "fmt"

func Sqrt(x float64)  float64 {
	var i int
	var z float64 = 1

	for i = 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}

	return z
}

func main() {
	fmt.Println("====================================================================")
	var number, sqrt float64
	fmt.Println("What number would you like to take the square root of?")
	fmt.Scanf("%f", &number)

	sqrt = Sqrt(number)

	fmt.Print("The square root of ", number, " is ", sqrt)
	fmt.Println(" ")
	fmt.Println("====================================================================")
}
