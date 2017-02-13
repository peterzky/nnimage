package main

import "github.com/peterzky/nnimage/base"
import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// genColor ...
func genColor(intr int) [][]uint8 {
	output := [][]uint8{}
	for i := 0; i <= 255; i += intr {
		for j := 0; j <= 255; j += intr {
			for k := 0; k <= 255; k += intr {
				a := []uint8{uint8(i), uint8(j), uint8(k)}
				output = append(output, a)

			}
		}
	}
	return output

}

func main() {
	intr, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	colors := genColor(int(intr))
	for _, v := range colors {
		base.Genimage(v)
		fmt.Println(v)
	}

}
