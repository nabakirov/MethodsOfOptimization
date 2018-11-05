package core


func BisectionSearchMethod(a, b, tol float64, f func(float64) float64) float64{
	var m float64
	fa, fb := f(a), f(b)
	sign := func(number float64) bool {return number < 0}

	if sign(fa) == sign(fb) {
		panic("Sign on f(a) and f(b) must be opposite!\nCheck endpoints of the interval [a, b]!")
	} else {
		k := 0
		k_max := 6
		for (fa - fb) > tol && k < k_max {
			k += 1
			m = a + (b - a) / 2
			fa := f(a)
			fm := f(m)

			if sign(fa) == sign(fm) {
				a = m
			} else {
				b = m
			}
		}
	}
	return m
}