package main

import (
	"fmt"

	"github.com/peterzky/nnimage/base"
)

// readImage to array
func readImage(path string) [][]float64 {
	p := base.ImageToPixels(path)
	pa := base.ConcatPixels(p)
	return pa
}

func main() {
	inputs := readImage("./img/test.png")
	// targets := readImage("./img/test2.png")
	// nn := base.NewNetwork(4, 5, 4, true, 0.1, 0.1)
	// nn.Train(inputs, targets, 1000)
	// base.DumpNN("./dump/image.json", nn)
	for _, v := range inputs {
		fmt.Println(v)
	}

}
