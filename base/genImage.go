package base

import (
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

// Create hex-decimal file name
func HexRename(c []uint8) string {
	return hex.EncodeToString(c)
}

// Create sequential file name
func SeqRename(i int) string {
	num := strconv.Itoa(i)
	return zero(4-len(num)) + num

}

func zero(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "0"
	}
	return result
}

// Generate image with specified color
func GenImage(c []uint8, n string) {
	img := image.NewRGBA(image.Rect(0, 0, 1024, 600))

	for i := 0; i <= 1024; i++ {
		for j := 0; j <= 600; j++ {
			img.Set(i, j, color.RGBA{c[0], c[1], c[2], 255})
		}
	}
	name := "./img/" + n + ".png"
	f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

}

// Set an interval (eg. 51) return a list of []uint8
func GenColor(intr int) [][]uint8 {
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
