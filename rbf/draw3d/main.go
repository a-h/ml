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

	anim := gif.GIF{LoopCount: 64}

	delay := 5 // 50ms

	deviationFrom, deviationTo := 0.0, 10.0
	deviationStep := (deviationFrom - deviationTo) / float64(anim.LoopCount)

	for i := 0; i < anim.LoopCount; i++ {
		img := image.NewPaletted(imgSize, palette)
		center := []float64{0.0, 0.0}
		deviation := []float64{deviationFrom + deviationStep, deviationFrom + deviationStep}
		gaussian, err := rbf.NewGaussianVector(1.0, center, deviation)
		if err != nil {
			fmt.Println("Failed to create Gaussian function:", err)
			return
		}
		f := func(x, y float64) float64 {
			return gaussian([]float64{x, y})
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
