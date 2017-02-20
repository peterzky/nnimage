package main

import "github.com/peterzky/nnimage/base"
import "fmt"

func main() {
	//5*3*17
	colors := base.GenColor(85)
	for i, v := range colors {
		name := base.SeqRename(i) + "-" + base.HexRename(v)
		base.GenImage(v, name)
		fmt.Println(v, name)
	}

}
