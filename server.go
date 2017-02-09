package main

import (
    "image/color"

    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/vg"
    "github.com/gonum/stat/distuv"
	"github.com/peterzky/nnimage/base"
)

func main() {
    // Draw some random values from the standard
    // normal distribution.

	m := base.ImageToPixels("./img/0091ea.png")
	p := base.ConcatPixels(m)
	var a []float64
	for _,v := range p {
		a = append(a,v[0])
	}

    v := make(plotter.Values, 10000)
    for i := range v {
        v[i] = a[i]
    }

    // Make a plot and set its title.
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Histogram"

    // Create a histogram of our values drawn
    // from the standard normal.
    h, err := plotter.NewHist(v, 16)
    if err != nil {
        panic(err)
    }
    // Normalize the area under the histogram to
    // sum to one.
    h.Normalize(1)
    p.Add(h)

    // The normal distribution function
    norm := plotter.NewFunction(distuv.UnitNormal.Prob)
    norm.Color = color.RGBA{R: 255, A: 255}
    norm.Width = vg.Points(2)
    p.Add(norm)

    // Save the plot to a PNG file.
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
        panic(err)
    }
}
