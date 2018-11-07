package main

import (
	"github.com/nabakirov/MethodsOfOptimization/core"
	_"github.com/nabakirov/MethodsOfOptimization/cli"
	"fmt"
	"math"
	_"github.com/tucnak/climax"
)

func f(x float64) float64 {
	return -0.04 * math.Pow(x, 3) + math.Pow(x, 2) + x - 1
}


func bisectionTest() {
	fmt.Println("Bisection Search Method")
	a := -1.6
	b := 0.6
	// a := -2
	// b := 0.8
	tol := 0.01
	result := core.BisectionSearchMethod(a, b, tol, 40, f)
	fmt.Println(result)
}


func evenSearchTest() {
	fmt.Println("Even Search Method")
	var h float64 = 0.2
	var x0 float64 = -3
	var k_max int = 10000000
	var tol float64 = 0.1

	result_min := core.EvenSearchMethodMin(h, x0, tol, k_max, f)
	result_max := core.EvenSearchMethodMax(h, x0, tol, k_max, f)

	fmt.Println("min =", result_min)
	fmt.Println("max =", result_max)
}

func pocketSearchTest() {
	fmt.Println("Pocket Search Method")
	var x0 float64 = 1
	var h0 float64 = 4
	var r float64 = 10
	var tol float64 = 0.01
	var k_max int = 10000

	result_min := core.PockerSearchMethodMin(x0, h0, r, tol, k_max, f)
	result_max := core.PockerSearchMethodMax(x0, h0, r, tol, k_max, f)

	fmt.Println("min", result_min)
	fmt.Println("max", result_max)
}

func goldenSectionSearchTest() {
	fmt.Println("Golden Section Search Method")
	var a float64 = -10
	var b float64 = 3
	var tol float64 = 0.01
	var k_max int = 10000

	result_min := core.GoldenSectionSearchMethodMin(a, b, tol, k_max, f)
	result_max := core.GoldenSectionSearchMethodMax(a, b, tol, k_max, f)

	fmt.Println("min", result_min)
	fmt.Println("max", result_max)

}



func main() {
	evenSearchTest()
	// mo := climax.New("mo")

	// mo.AddCommand(cli.EsmCliCommand)
	// mo.Run()
}