package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type biColor struct {
	left, right []uint8
}

// color create color template
func gencolor(intr1, intr2 int) []biColor {
	var result []biColor
	for i := 0; i <= 255; i += intr1 {
		for j := -45; j <= 45; j += intr2 {
			if i+j >= 0 && i+j <= 255 && i != i+j {
				left := []uint8{uint8(i), 0, 0}
				right := []uint8{uint8(i + j), 0, 0}
				color2 := biColor{left, right}
				result = append(result, color2)
				fmt.Printf("%v + %v = %v %+v\n", i, j, i+j, color2)
			}
		}
	}
	return result
}

// create image base on biColor input
func genImage2(c biColor, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 100))
	l := c.left
	r := c.right
	for i := 0; i < 100; i++ {
		for j := 0; j <= 100; j++ {
			img.Set(i, j, color.RGBA{l[0], l[1], l[2], 255})
		}
	}
	for i := 100; i < 200; i++ {
		for j := 0; j <= 100; j++ {
			img.Set(i, j, color.RGBA{r[0], r[1], r[2], 255})
		}
	}
	name := "./img/" + filename + ".png"
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic("fail to open file")
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		panic("error encode png")
	}

}

func main() {
	colors := gencolor(15, 15)
	for i, c := range colors {
		genImage2(c, strconv.Itoa(i))
	}
}
