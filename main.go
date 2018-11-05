package main

import (
	"github.com/nabakirov/mo/core"
	"fmt"
	"math"
)


func main() {
	f := func(x float64) float64 {return -0.04 * math.Pow(x, 3) + math.Pow(x, 2) + x - 1}
	a := -1.6
	b := 0.6
	// a := -2
	// b := 0.8
	tol := 0.01
	result := core.BisectionSearchMethod(a, b, tol, f)
	fmt.Println(result)
}