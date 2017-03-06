package main

import "fmt"
import "io"
import "image"

type biColor struct {
	left, right []uint8
}

// color create color template
func color(intr1, intr2 int) []biColor {
	var result []biColor
	for i := 0; i <= 255; i += intr1 {
		for j := -45; j <= 45; j += intr2 {
			if i+j >= 0 && i+j <= 255 {
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
func genImage2(c biColor, w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	l := c.left
	r := c.right
	for i := 0; i < 100; i += 1 {
		for j := 0; j <= 100; j += 1 {
			img.Set(i, j, color.RGBA{l[0], l[1], l[2], 255})
		}
	}
	for i := 100; i < 200; i += 1 {
		for j := 0; j <= 100; j += 1 {
			img.Set(i, j, color.RGBA{r[0], r[1], r[2], 255})
		}
	}
	w.Write(img)

}

func main() {
	colors := color(15, 15)

	for _, c := range colors {
		for i := 0; i <= 255; i += 15 {
			if c.left[0] == uint8(i) {
				fmt.Println(c)
			}
		}
	}
}
