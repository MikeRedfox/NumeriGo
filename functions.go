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
