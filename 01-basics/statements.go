package main

import "fmt"

func loopToN(n int) {
	fmt.Printf("1 to %v:\n", n)
	for i := 1; i < n+1; i++ {
		fmt.Print(i, " ")
	}
}

func PrintLoops(n int) {

	loopToN(n)
}
