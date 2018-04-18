package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gonum.org/v1/plot/plotter"

	"github.com/a-h/ml/rbf"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	p, err := plot.New()
	if err != nil {
		fmt.Println("Error creating Plot: ", err)
		os.Exit(-1)
	}

	p.Title.Text = "Gaussian and Ricker Wavelet RBFs"
	p.X.Min = 0
	p.X.Padding = 0
	p.X.Label.Text = "X"
	p.Y.Min = 0
	p.Y.Padding = 0
	p.Y.Label.Text = "Y"

	g := rbf.NewGaussian(1.0, 0.0, 0.5)
	rw := rbf.NewRickerWavelet(1.0, 0.0, 1.0)

	var gaussian, rickerwavelet plotter.XYs
	step := 0.1
	for i := -6.0; i < 6.0; i += step {
		gaussian = append(gaussian, xy{
			X: float64(i),
			Y: g(float64(i)),
		})
		rickerwavelet = append(rickerwavelet, xy{
			X: float64(i),
			Y: rw(float64(i)),
		})
	}

	plotutil.AddLines(p, gaussian, rickerwavelet)

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, "points.png"); err != nil {
		panic(err)
	}
}

type xy struct {
	X, Y float64
}
