package main

import "github.com/peterzky/nnimage/base"

func main() {
	s := base.GenSamples("./img", "png")
	in, ta := base.GenData(s)
	nn := base.NewNetwork(3, 4, 3, true, 0.1, 0.1)
	nn.Train(in, ta, 5)
	base.DumpNN("dump.json", nn)
}
