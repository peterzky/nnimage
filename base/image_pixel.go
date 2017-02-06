package base

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// ConcatPixels convert to array
func ConcatPixels(p [][]Pixel) [][]float64 {
	b := []Pixel{}
	f := [][]float64{}
	for _, v := range p {
		b = append(b, v...)
	}
	for _, x := range b {
		ff := []float64{x.R, x.G, x.B}
		f = append(f, ff)
	}
	return f

}

// convert image to pixel array
func ImageToPixels(path string) [][]Pixel {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return pixels
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{float64(r / 257), float64(g / 257), float64(b / 257), float64(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R float64
	G float64
	B float64
	A float64
}
