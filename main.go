package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./img.png")
	defer file.Close()

	disposition := [][]int{{-5, 0}, {10, 5}, {0, 7}}

	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()

	wImg := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x+disposition[0][0], y+disposition[0][1]).RGBA()
			_, g, _, _ := img.At(x+disposition[1][0], y+disposition[1][1]).RGBA()
			_, _, b, _ := img.At(x+disposition[2][0], y+disposition[2][1]).RGBA()
			_, _, _, a := img.At(x, y).RGBA()
			newColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
			wImg.SetRGBA(x, y, newColor)
		}
	}

	result, err := os.Create("./res.png")
	defer result.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(result, wImg)
	if err != nil {
		log.Fatal(err)
	}

}
