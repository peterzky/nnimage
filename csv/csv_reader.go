package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/peterzky/nnimage/base"
)

//Parse CSV data into Sample sets for neural network
func ParseCSV(filename string) []base.Sample {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to load")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	var sampleSets []base.Sample

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			panic("Parse error")
		}

		var colorfloat []float64

		for _, color := range record {
			f, err := strconv.ParseFloat(color, 64)
			if err != nil {
				panic("Can't convert to float64")
			}
			colorfloat = append(colorfloat, f)
		}

		var sample base.Sample
		sample.Input = colorfloat[:3]
		sample.Target = colorfloat[3:]
		sampleSets = append(sampleSets, sample)
	}
	return sampleSets
}

func main() {
	sampleSets := ParseCSV("color-map.csv")
	fmt.Println(sampleSets)
	input, target := base.GenData(sampleSets)
	nn := base.NewNetwork(3, 50, 3, true, 0.0000001, 0.0000001)
	nn.Train(input, target, 5000000)
	base.DumpNN("dumnp.jsonb", nn)

}
