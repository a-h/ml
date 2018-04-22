package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"

	"github.com/a-h/ml/projection"
	"github.com/a-h/ml/rbf"
)

var palette = []color.Color{color.White, color.Black}

func main() {
	file, err := os.Create("op.gif")
	if err != nil {
		fmt.Println("Failed to create op.gif file:", err)
		return
	}

	imgSize := image.Rect(0, 0, 500, 500)

	anim := gif.GIF{LoopCount: 128}

	delay := 1 // 10ms

	deviationFrom, deviationTo := 0.0, 10.0
	deviationStep := (deviationFrom - deviationTo) / float64(anim.LoopCount)

	for i := 0; i < anim.LoopCount; i++ {
		img := image.NewPaletted(imgSize, palette)
		center := []float64{0.0, 0.0}
		deviation := deviationFrom + deviationStep
		gaussian := rbf.NewGaussianVector(1.0, center, deviation)
		f := func(x, y float64) (z float64) {
			z, err := gaussian([]float64{x, y})
			if err != nil {
				panic(err)
			}
			return z
		}
		projectionAngle := 30.0
		d := projection.New(-1.0, 1.0, f, imgSize.Dx(), imgSize.Dy(), projectionAngle)
		d.Draw(img)

		// Add the image to the output.
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		deviationFrom += deviationStep
	}

	gif.EncodeAll(file, &anim)
}
