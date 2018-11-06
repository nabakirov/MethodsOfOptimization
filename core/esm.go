package core

import _"fmt"


func EvenSearchMethodMin(h, x0, tol float64, k_max int, f func(float64) float64) float64{
	k := 0
	yf0 := f(x0)
	x1 := x0 + h
	yf1 := f(x1)

	for k < k_max{
		k += 1
		if yf1 <= yf0{
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 += h
			yf1 = f(x1)
		}
	}
	return yf1
}

func EvenSearchMethodMax(h, x0, tol float64, k_max int, f func(float64) float64) float64{
	k := 0
	yf0 := f(x0)
	x1 := x0 + h
	yf1 := f(x1)

	for k < k_max{
		k += 1
		if yf1 >= yf0{
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 += h
			yf1 = f(x1)
		}
	}
	return yf1
}