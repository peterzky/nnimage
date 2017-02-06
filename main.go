package main

import (
	"fmt"

	"github.com/peterzky/nnimage/base"
)

func main() {
	s := base.GenSamples("./img")
	in, ta := base.Data(s)
	for i := 0; i < 10; i++ {
		fmt.Println(in[i], ta[i])
	}
}
