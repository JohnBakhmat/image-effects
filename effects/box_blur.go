package effects

import (
	"image"
	"image/color"
	"log"
	"math"
)

func average(window []color.Color) color.Color {
	red, green, blue, alpha := 0.0, 0.0, 0.0, 0.0
	for _, c := range window {
		r, g, b, a := c.RGBA()
		red += float64(r)
		green += float64(g)
		blue += float64(b)
		alpha += float64(a)
	}

	return color.RGBA{
		R: uint8(math.Round(red / float64(len(window)))),
		G: uint8(math.Round(green / float64(len(window)))),
		B: uint8(math.Round(blue / float64(len(window)))),
		A: uint8(math.Round(alpha / float64(len(window)))),
	}

}

func BoxBlur(img image.Image, kernel int) image.Image {
	bounds := img.Bounds()
	result_image := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	area := int((kernel*2 + 1) * (kernel*2 + 1))
	for y := bounds.Min.Y + kernel; y < bounds.Max.Y-kernel; y++ {
		for x := bounds.Min.X + kernel; x < bounds.Max.X-kernel; x++ {

			window := make([]color.Color, 0, area)

			for i := x - kernel; i <= x+kernel; i++ {
				for j := y - kernel; j <= y+kernel; j++ {
					window = append(window, img.At(i, j))
				}
			}

			result_image.Set(x, y, average(window))
		}
	}
	return result_image
}
