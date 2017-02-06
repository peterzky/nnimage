package base

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Sample struct {
	Input, Target []float64
}

// generate inputs and targets for training
func GenData(s []Sample) (inputs, targets [][]float64) {
	for _, v := range s {
		inputs = append(inputs, v.Input)
		targets = append(targets, v.Target)
	}
	return inputs, targets
}

// GenTarget teacher generator
func genTarget(h string, l int) [][]float64 {
	target := [][]float64{}
	a, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	var b []float64
	for _, v := range a {
		b = append(b, float64(v))
	}
	for i := 0; i < l; i++ {
		target = append(target, b)
	}
	return target

}

// shuffle smaples order
func shuffle(s []Sample) []Sample {
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s

}

// read image to pixel array
func readImage(path string) [][]float64 {
	p := ImageToPixels(path)
	pa := ConcatPixels(p)
	return pa
}

// generator target pair with input string
func genTrain(path string, h string) []Sample {
	inputs := readImage(path)
	targets := genTarget(h, len(inputs))
	var samples []Sample
	for i := 0; i < len(inputs); i++ {
		s := Sample{Input: inputs[i], Target: targets[i]}
		samples = append(samples, s)
	}
	return samples
}

func GenSamples(p, ext string) []Sample {
	var s []Sample
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		path := p + "/" + file.Name()
		hex := strings.TrimSuffix(file.Name(), "."+ext)
		x := genTrain(path, hex)
		s = append(s, x...)
	}
	return shuffle(s)

}
