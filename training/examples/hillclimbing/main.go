package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"

	"github.com/a-h/ml/training"

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

	anim := gif.GIF{LoopCount: 32}

	delay := 1 // 10ms

	// Make some "terrain" out of 3D gaussian functions.
	hills := []rbf.VectorFunction{
		rbf.NewGaussianVector(1.0, []float64{-0.5, -0.25}, 0.25),
		rbf.NewGaussianVector(0.75, []float64{0.25, -0.25}, 0.5),
		rbf.NewGaussianVector(0.5, []float64{-0.25, 0.25}, 0.4),
		rbf.NewGaussianVector(0.8, []float64{0.5, 0.25}, 0.3),
	}

	f := func(x, y float64) (z float64) {
		for _, f := range hills {
			r, err := f([]float64{x, y})
			if err != nil {
				panic(err)
			}
			if r > z {
				z = r
			}
		}
		return z
	}

	trainee := FuncTrainee{
		F: f,
	}

	highest := math.MaxFloat64
	evaluator := func() (e float64, err error) {
		z := f(trainee.X, trainee.Y)
		return highest - z, nil
	}

	// Start off the trainee at a location.
	memory := []float64{1, 1}
	t := training.NewHillClimbing(memory, 0.2, 0.1)

	// Keep walking.
	for i := 0; i < anim.LoopCount; i++ {
		img := image.NewPaletted(imgSize, palette)
		projectionAngle := 30.0

		vis := func(x, y float64) (z float64) {
			z = f(x, y)
			// Put a steep hill where the trainee hill climber currently is...
			pos, err := rbf.NewGaussianVector(1.0, []float64{trainee.X, trainee.Y}, 0.05)([]float64{x, y})
			if err != nil {
				panic(err)
			}
			if pos > z {
				z = z + (pos * 0.1) // 10% higher
			}
			return z
		}

		d := projection.New(-1.0, 1.0, vis, imgSize.Dx(), imgSize.Dy(), projectionAngle)
		d.Draw(img)

		// Add the image to the output.
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		// Move some more.
		memory, _ := t.Next(evaluator)
		fmt.Println(memory)
		trainee.SetMemory(memory)
	}

	gif.EncodeAll(file, &anim)
}

type FuncTrainee struct {
	F    func(x, y float64) (z float64)
	X, Y float64
}

func (ft *FuncTrainee) Calculate(input []float64) (output []float64, err error) {
	x, y := input[0], input[1]
	z := ft.F(x, y)
	return []float64{z}, nil
}

func (ft *FuncTrainee) GetMemorySize() int {
	return 2
}

func (ft *FuncTrainee) GetMemory() []float64 {
	return []float64{ft.X, ft.Y}
}

func (ft *FuncTrainee) SetMemory(m []float64) {
	ft.X = m[0]
	ft.Y = m[1]
}
