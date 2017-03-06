package main

import "github.com/peterzky/nnimage/base"
import "github.com/codahale/hdrhistogram"
import "fmt"

func main() {
	m := base.ImageToPixels("./sample/00ff00.png")
	p := base.ConcatPixels(m)
	var a []int64
	for _, v := range p {
		a = append(a, int64(v[2]))
	}
	hist := hdrhistogram.New(0, 255, 1)
	for _, sample := range a {
		hist.RecordValue(sample)
	}
	fmt.Println(hist.Mean())
	fmt.Println(hist.Distribution())
	fmt.Println(len(a))
	// cod := []string{"b7", "c6", "d3", "e5", "f4", "ef", "e5", "d5", "ff", "88"}
	// for _, v := range cod {
	// 	fmt.Println(hex.DecodeString(v))
	// }

}
