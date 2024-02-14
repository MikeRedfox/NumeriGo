package numerigo

import (
	"errors"
)

func PuntoFisso[T int | int8 | int16 | int32 | int64 | float32 | float64](x0, toll T, f func(T) T, nmax int) (T, error) {

	var diff T = 1.0
	n := 1
	var xs [100]T
	xs[0] = x0

	for diff > toll && n < nmax {
		xs[n] = f(xs[n-1])
		diff = xs[n] - xs[n-1]
		if diff < 0 {
			diff = diff * -1
		}
		n++
	}

	if n > nmax {
		return 0, errors.New("Reached number of max iteration, Seggesting to change x0 value")
	}

	return xs[n-1], nil

}

func IsEven[T int | int16 | int32 | int64](n T) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func Simpson[T int | int8 | int16 | int32 | int64 | float32 | float64](f func(float64) float64, a, b T, m int) float64 {
	if !IsEven(m) {
		m++
	}

	var I float64

	h := float64((b - a)) / float64(m)
	I = f(float64(a)) + f(float64(b))

	for i := 1; i <= m-1; i += 2 {
		x := float64(a) + float64(i)*h
		I = I + 4*f(x)
	}

	for i := 2; i < m-2; i += 2 {
		x := float64(a) + float64(i)*h
		I = I + 4*f(x)
	}
	I = float64(h) * float64(I) / float64(3)
	return I

}
