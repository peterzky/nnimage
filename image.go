package main

import "./lib"

// readImage to array
func readImage(path string) []float64 {
	p := lib.ImageToPixels(path)
	pa := lib.ConcatPixels(p)
	return pa
}

func main() {
	inputs := readImage("./img/test.png")
	targets := readImage("./img/test2.png")
	nn := lib.NewNetwork(4, 5, 4, true, 0.1, 0.1)
	nn.Train(inputs, targets, 1000)
	nn.DumpNN("./dump/image.json")

}
