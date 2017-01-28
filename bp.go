package main

import (
	"fmt"
	"./gonn"
)

func main(){
	nn := gonn.NewNetwork(2,3,1,true,0.0001,0.01)
	//定义输入样本 
	inputs := [][]float64 {
		[]float64{0,0},
		[]float64{0,1},
		[]float64{1,0},
		[]float64{1,1},

		[]float64{2,3},
		[]float64{3,4},
		[]float64{4,5},
		[]float64{5,6},


		[]float64{10,11},
		[]float64{11,12},
		[]float64{12,13},
		[]float64{13,14},

	}

	targets := [][]float64 {
		[]float64{0},//0+0=0
		[]float64{1},//0+1=1
		[]float64{1},//1+0=1
		[]float64{2},//1+1=2

		[]float64{5},
		[]float64{7},
		[]float64{9},
		[]float64{11},


		[]float64{21},
		[]float64{23},
		[]float64{25},
		[]float64{27},
	}

	// for i := 0 ; i < 1000 ; i++ {
	// 	inputs = append(inputs,[]float64{float64(i)})
	// 	targets = append(targets,[]float64{float64(i+2)})
	// }

	// fmt.Println(inputs,targets)

	

	nn.Train(inputs,targets,1000000)

	for _,p := range inputs{
		fmt.Println(nn.Forward(p))
	}

	// testcase := []float64{11,12}

	// test := nn.Forward(testcase)

	// fmt.Println(test)
	gonn.DumpNN("dump.json",nn)

}
