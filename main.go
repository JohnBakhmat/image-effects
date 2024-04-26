package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/johnbakhmat/image-effects/effects"
)

func read_image(path string) image.Image {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
func write_image(path string, img image.Image) {
	result, err := os.Create(path)
	defer result.Close()

	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(result, img)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img := read_image("./img.png")
	ca_img := effects.ChromaticAberration(effects.ChromaticAberrationProps{
		Red_displacement:   [2]int{-20, 0},
		Green_displacement: [2]int{40, 15},
		Blue_displacement:  [2]int{0, 27},
		Img:                img,
		Strength:           0.5,
	})

	write_image("./ca_res.png", ca_img)

}
