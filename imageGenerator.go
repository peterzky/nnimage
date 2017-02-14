package main

import "github.com/peterzky/nnimage/base"
import "fmt"

func main() {
	// intr, err := strconv.ParseInt(os.Args[1], 10, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	colors := base.GenColor(51)
	for i, v := range colors {
		base.GenImage(v, base.SeqRename(i))
		fmt.Println(v, i)
	}

}
