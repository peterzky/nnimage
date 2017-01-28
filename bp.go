package main

import (
	"fmt"
	"./gonn"
)

func main(){
	nn := gonn.NewNetwork(3,4,3,true,0.000001,0.0000001)
	inputs := make([][]float64, 15)
	targets := make([][]float64,15)

	for i := 0; i < len(inputs);i++ {
		x := float64(i)
		inputs[i] = []float64{x,x+1,x+2}
		targets[i] = []float64{x*(x+1),(x+1)*(x+2),x+3}
	}
	fmt.Println(inputs,targets)

	

	nn.Train(inputs,targets,10000000)

	for _,p := range inputs{
		fmt.Println(nn.Forward(p))
	}

	// testcase := []float64{11,12}

	// test := nn.Forward(testcase)

	// fmt.Println(test)
	gonn.DumpNN("dump.json",nn)

}
