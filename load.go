package main

import (
	"fmt"
	"os"
	"strconv"

	gonn "./lib"
)

func main() {
	var numbers []float64
	nn := gonn.LoadNN("./dump/dump.json")
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
