package base

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/peterzky/nnimage/base"
)

// crop x1,y1 is the anchor. x2,y2 is the size
func crop(src image.Image, x1, y1, x2, y2 int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, x2, y2))

	for i := 0; i <= x2; i++ {
		for j := 0; j <= y2; j++ {
			img.Set(i, j, src.At(i+x1, j+y1))
		}
	}
	return img

}

// Crop image with default position
func CropImg(path, name string) {
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	out := crop(img, 200, 200, 100, 100)

	sample := "./sample/" + name + ".png"

	f, _ := os.OpenFile(sample, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	e := png.Encode(f, out)
	if e != nil {
		log.Fatal(e)
	}

}

func CreateCropSample(path string) {
	path := "/home/peterzky/Downloads/downloads/一数科技/screenshots/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	colors := base.GenColor(85)
	var hex []string
	for _, c := range colors {
		hex = append(hex, base.HexRename(c))

	}
	hex = hex[1:]
	for i, file := range files {
		f := path + file.Name()
		CropImg(f, hex[i])
	}
}
