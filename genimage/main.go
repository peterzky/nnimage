package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type biColor struct {
	left, right    []uint8
	labell, labelr string
}

// filename for current colors
func (c *biColor) filename() string {
	return c.labell + "-" + c.labelr

}

// addLabel addlabel to image
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

// dispatcher dispatch rgb specification
func rgbDispatcher(rgb string, i, j int) biColor {
	var result biColor
	switch rgb {
	case "R":
		result.left = []uint8{uint8(i), 0, 0}
		result.right = []uint8{uint8(i + j), 0, 0}
		result.labell = "R:" + strconv.Itoa(i)
		result.labelr = "R:" + strconv.Itoa(i+j)
		return result
	case "G":
		result.left = []uint8{0, uint8(i), 0}
		result.right = []uint8{0, uint8(i + j), 0}
		result.labell = "G:" + strconv.Itoa(i)
		result.labelr = "G:" + strconv.Itoa(i+j)
		return result
	case "B":
		result.left = []uint8{0, 0, uint8(i)}
		result.right = []uint8{0, 0, uint8(i + j)}
		result.labell = "B:" + strconv.Itoa(i)
		result.labelr = "B:" + strconv.Itoa(i+j)
		return result
	default:
		panic("Invalid RGB Notation")
	}

}

// color create color template
func gencolor(intr1, intr2 int, c string) []biColor {
	var result []biColor
	for i := 0; i <= 255; i += intr1 {
		for j := -45; j <= 45; j += intr2 {
			if i+j >= 0 && i+j <= 255 && i != i+j {
				color2 := rgbDispatcher(c, i, j)
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

	addLabel(img, 30, 20, c.labell)
	addLabel(img, 130, 20, c.labelr)
	name := "./img/" + filename + ".png"
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic("fail to open file")
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

}

func main() {
	colorsR := gencolor(15, 15, "R")
	colorsG := gencolor(15, 15, "G")
	colorsB := gencolor(15, 15, "B")
	colors := append(append(colorsR, colorsG...), colorsB...)
	for _, c := range colors {
		genImage2(c, c.filename())
	}
}
