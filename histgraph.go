package main

import "github.com/peterzky/nnimage/base"
import "encoding/csv"
import "os"
import "log"
import "strconv"

func parsePixel(p [][]float64) (out [][]string) {
	for _, x := range p {
		r := strconv.FormatFloat(x[0], 'f', 0, 64)
		g := strconv.FormatFloat(x[1], 'f', 0, 64)
		b := strconv.FormatFloat(x[2], 'f', 0, 64)
		out = append(out, []string{r, g, b})
	}
	return out

}

func main() {
	m := base.ImageToPixels("/home/peterzky/Downloads/downloads/一数科技/手/hand2.jpg")
	records := parsePixel(base.ConcatPixels(m))

	out_f, err := os.OpenFile("distro.csv", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("error!")
	}

	defer out_f.Close()
	w := csv.NewWriter(out_f)

	for i, record := range records {
		if i < 1000 {
			if err := w.Write(record); err != nil {
				log.Fatal("Error writing data")
			}
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
