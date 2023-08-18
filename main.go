package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/cenkalti/dominantcolor"
)

func FindDominantColors(fileInput string) ([]color.RGBA, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		fmt.Println("File not found:", fileInput)
		return []color.RGBA{}, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return []color.RGBA{}, err
	}

	colors := dominantcolor.FindN(img, 3)
	return colors, nil
}

func main() {
	colors, err := FindDominantColors("cat.jpeg")
	if err != nil {
		panic(err)
	}

	for n, c := range colors {
		fmt.Printf("Color %d: %s https://www.color-hex.com/color/%s\n", n+1, dominantcolor.Hex(c), dominantcolor.Hex(c)[1:])
	}
}
