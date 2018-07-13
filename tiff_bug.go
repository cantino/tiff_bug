package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"

	"golang.org/x/image/tiff"
)

func main() {
	width, height := 4104, 2736

	i := image.NewRGBA64(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i.Set(x, y, color.RGBA64{R: math.MaxUint16, G: math.MaxUint16, B: math.MaxUint16, A: math.MaxUint16})
		}
	}

	outputFile, err := os.Create("out-compressed.tiff")
	if err != nil {
		log.Fatal(err)
	}
	tiff.Encode(outputFile, i, &tiff.Options{Compression: tiff.Deflate})

	outputFile, err = os.Create("out-noncompressed.tiff")
	if err != nil {
		log.Fatal(err)
	}
	tiff.Encode(outputFile, i, &tiff.Options{})
}
