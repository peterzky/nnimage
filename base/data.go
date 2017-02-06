package base

import "encoding/hex"

// GenTarget teacher generator
func GenTarget(h string, l int) [][]float64 {
	target := [][]float64{}
	a, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	var b []float64
	for _, v := range a {
		b = append(b, float64(v))
	}
	for i := 0; i < l; i++ {
		target = append(target, b)
	}
	return target

}
