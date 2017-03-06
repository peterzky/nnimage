package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"github.com/peterzky/nnimage/base"
)

// getRGB return r,g,b array base on parameter
func getRGB(a [][]float64, i int) []float64 {
	var result []float64
	for _, p := range a {
		result = append(result, p[i])
	}
	return result
}

// average return average from an array
func average(xs []float64) float64 {
	var total float64
	for _, v := range xs {
		total += v
	}
	return math.Ceil(total / float64(len(xs)))
}

func genavg(p, ext string) []base.Sample {
	var result []base.Sample
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		path := p + "/" + file.Name()
		m := base.ImageToPixels(path)
		pixels := base.ConcatPixels(m)
		r := getRGB(pixels, 0)
		g := getRGB(pixels, 1)
		b := getRGB(pixels, 2)
		avg := []float64{average(r), average(g), average(b)}
		real, err := hex.DecodeString(strings.TrimSuffix(file.Name(), "."+ext))
		var target []float64
		for _, v := range real {
			target = append(target, float64(v))
		}
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, base.Sample{avg, target})
	}
	return result

}

func main() {
	avg := genavg("./sample", "png")
	var slice []base.Sample
	for _, v := range avg {
		if v.Target[0] == 0 {
			slice = append(slice, v)
		}

	}
	// for _, v := range slice {
	// 	fmt.Println(v)
	// }
	// // fmt.Println(len(slice))
	in, ta := base.GenData(slice)
	nn := base.NewNetwork(3, 60, 3, true, 0.0000001, 0.0000001)

	nn.Train(in, ta, 30000000)
	// // reverse
	// nn.Train(ta, in, 90000)
	base.DumpNN("dump.json", nn)

}
