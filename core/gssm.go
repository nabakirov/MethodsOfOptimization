package core
import "math"


func GoldenSectionSearchMethodMin(a, b, tol float64, k_max int, f func(float64) float64) float64 {
	r := (math.Sqrt(5) - 1) / 2
	x1 := a + (1 - r) * (b - a)
	f1 := f(x1)
	x2 := a + r * (b - a)
	f2 := f(x2)
	k := 0

	for k < k_max {
		k += 1
		if f1 > f2 {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + r * (b - a)
			f2 = f(x2)
		} else {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + (1 - r) * (b - a)
			f1 = f(x1)
		}
	}
	return f2
}
func GoldenSectionSearchMethodMax(a, b, tol float64, k_max int, f func(float64) float64) float64 {
	r := (math.Sqrt(5) - 1) / 2
	x1 := a + (1 - r) * (b - a)
	f1 := f(x1)
	x2 := a + r * (b - a)
	f2 := f(x2)
	k := 0

	for k < k_max {
		k += 1
		if f1 < f2 {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + r * (b - a)
			f2 = f(x2)
		} else {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + (1 - r) * (b - a)
			f1 = f(x1)
		}
	}
	return f2
}