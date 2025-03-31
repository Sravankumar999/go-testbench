package main

import (
	"fmt"
	"math"
	"math/rand"

	"rsc.io/quote"
)

func swap(x, y string) (a, b string) {
	// named returns
	a = y
	b = x
	return
}

func PrintFmts() {

	fmt.Println("Init")
	fmt.Println(quote.Go())
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(5))
	a, b := "myself", "hi "
	fmt.Println(swap(a, b))

}

func PrintSqRoot(x float64) {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	fmt.Println(z)
}
