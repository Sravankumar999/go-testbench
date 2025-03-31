package main

import "fmt"

func main() {
	// Calling from fmt_funcs.go
	fmt.Println("=============== fmt:")
	PrintFmts()
	fmt.Println("=============== Loops:")
	PrintLoops(7)
	fmt.Println("=============== Sq root (Newton Formula):")
	PrintSqRoot(143)

}
