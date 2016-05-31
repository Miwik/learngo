package main

import "fmt"

func print(args ...interface{}) {
	for i, a := range args {
		fmt.Printf("#%02d (%T) %v\n", i, a, a)
	}
}

func main() {
	a := "initial"
	b, c := 1, 2
	d, e, f := false, "plop", 1.25
	print(a, b, c, d, e, f)
}
