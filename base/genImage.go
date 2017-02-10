package base

import (
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Genimage color
func Genimage(c []uint8) {
	img := image.NewRGBA(image.Rect(0, 0, 1024, 600))

	for i := 0; i <= 1024; i++ {
		for j := 0; j <= 600; j++ {
			img.Set(i, j, color.RGBA{c[0], c[1], c[2], 255})
		}
	}

	str := hex.EncodeToString(c)

	name := "./img/" + str + ".png"

	f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}
