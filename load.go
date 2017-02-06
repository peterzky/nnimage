package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/peterzky/nnimage/base"
)

func main() {
	var numbers []float64
	nn := base.LoadNN("dump.json")
	if len(os.Args) <= 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	fmt.Println(nn.Forward(numbers))

}
