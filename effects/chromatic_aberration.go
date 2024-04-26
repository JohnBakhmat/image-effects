package effects

import (
	"image"
	"image/color"
	"math"
)

type ChromaticAberrationProps struct {
	Red_displacement   [2]int
	Green_displacement [2]int
	Blue_displacement  [2]int
	Img                image.Image
	Strength           float32
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Abs(math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)))
}

func ChromaticAberration(props ChromaticAberrationProps) image.Image {
	bounds := props.Img.Bounds()
	result_image := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	r := props.Red_displacement
	g := props.Green_displacement
	b := props.Blue_displacement

	centerX := (bounds.Max.X + bounds.Min.X) / 2
	centerY := (bounds.Max.Y + bounds.Min.Y) / 2
	maxDistance := distance(float64(centerX), float64(centerY), float64(bounds.Min.X), float64(bounds.Min.Y))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			distance := distance(float64(centerX), float64(centerY), float64(x), float64(y))
			strength := distance / maxDistance * float64(props.Strength)

			rX := int(math.Round(float64(r[0]) * strength))
			rY := int(math.Round(float64(r[1]) * strength))
			gX := int(math.Round(float64(g[0]) * strength))
			gY := int(math.Round(float64(g[1]) * strength))
			bX := int(math.Round(float64(b[0]) * strength))
			bY := int(math.Round(float64(b[1]) * strength))

			r, _, _, _ := props.Img.At(x+rX, y+rY).RGBA()
			_, g, _, _ := props.Img.At(x+gX, y+gY).RGBA()
			_, _, b, _ := props.Img.At(x+bX, y+bY).RGBA()
			_, _, _, a := props.Img.At(x, y).RGBA()
			newColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
			result_image.SetRGBA(x, y, newColor)
		}
	}

	return result_image
}
