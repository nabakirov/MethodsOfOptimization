package core
import "math"

func PockerSearchMethodMin(x0, h0, r, tol float64, k_max int, f func(float64) float64) float64 {
	yf0 := f(x0)
	h1 := h0
	x1 := x0 + h1
	yf1 := f(x1)
	k := 0

	for k < k_max {
		k += 1

		if yf1 >= yf0 {
			if math.Abs(h0) < tol / r {
				h1 = h0
				x1 = x0
				yf1 = yf0
			} else {
				h1 = -h0 / r
				h0 = h1
				x0 = x1
				yf0 = yf1
				x1 = x0 + h1

				yf1 = f(x1)
			}
		} else {
			h1 = h0
			x0 = x1
			yf0 = yf1
			x1 = x0 + h1

			yf1 = f(x1)
		}
	}
	return yf1

}
func PockerSearchMethodMax(x0, h0, r, tol float64, k_max int, f func(float64) float64) float64 {
	yf0 := f(x0)
	h1 := h0
	x1 := x0 + h1
	yf1 := f(x1)
	k := 0

	for k < k_max {
		k += 1

		if yf1 <= yf0 {
			if math.Abs(h0) < tol / r {
				h1 = h0
				x1 = x0
				yf1 = yf0
			} else {
				h1 = -h0 / r
				h0 = h1
				x0 = x1
				yf0 = yf1
				x1 = x0 + h1

				yf1 = f(x1)
			}
		} else {
			h1 = h0
			x0 = x1
			yf0 = yf1
			x1 = x0 + h1

			yf1 = f(x1)
		}
	}
	return yf1

}