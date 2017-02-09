package main

import "image"
import "image/color"
import "image/png"
import "os"

func main() {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 1024, 600))

	for i := 0; i <= 1024; i++ {
		for j := 0; j <= 600; j++ {
			img.Set(i, j, color.RGBA{0, 255, 0, 255})
		}
	}

	// Draw a red dot at (2, 3)
	// img.Set(2, 3, color.RGBA{255, 0, 0, 255})

	// Save to out.png
	f, _ := os.OpenFile("./img/out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
