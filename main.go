package main

import "github.com/peterzky/nnimage/base"

func main() {
	s := base.GenSamples("./sample", "png")
	in, ta := base.GenData(s)
	nn := base.NewNetwork(3, 4, 3, true, 0.01, 0.001)
	nn.Train(in, ta, 100)
	base.DumpNN("dump.json", nn)
}
