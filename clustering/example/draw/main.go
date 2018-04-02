package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/a-h/ml/clustering"
	"github.com/a-h/ml/distance"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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

	p.Title.Text = "KMeans"
	p.X.Min = 0
	p.X.Padding = 0
	p.X.Label.Text = "X"
	p.Y.Min = 0
	p.Y.Padding = 0
	p.Y.Label.Text = "Y"

	// Create some random data and assign to n clusters.
	data := random2DVectors(50)
	n := 3
	assignment, err := clustering.KMeans(data, n, distance.Euclidean)
	if err != nil {
		fmt.Println("Error clustering data: ", err)
		os.Exit(-1)
	}

	// Get the clusters.
	clusters, err := clustering.Assign(data, assignment)
	if err != nil {
		fmt.Println("Error assigning data to clusters: ", err)
		os.Exit(-1)
	}

	// Convert them to scatter inputs (something that implements the XYer interface).
	for i, cluster := range clusters {
		scatter := convert2DVectorToPlotterXY(cluster)
		// Add them to the chart.
		err = addScatters(p, i, strconv.Itoa(i), scatter)
		if err != nil {
			panic(err)
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, "points.png"); err != nil {
		panic(err)
	}
}

func convert2DVectorToPlotterXY(v []clustering.Vector) plotter.XYs {
	pts := make(plotter.XYs, len(v))
	for i := 0; i < len(v); i++ {
		pts[i] = xy{
			X: v[i][0],
			Y: v[i][1],
		}
	}
	return pts
}

type xy struct {
	X, Y float64
}

func random2DVectors(n int) []clustering.Vector {
	op := make([]clustering.Vector, n)
	for i := 0; i < n; i++ {
		v := make(clustering.Vector, 2)
		randomise(v, -10, 10)
		op[i] = v
	}
	return op
}

func randomise(v []float64, min, max int) {
	for i := 0; i < len(v); i++ {
		v[i] = float64(rand.Intn(max-min) + min)
	}
}

func addScatters(plt *plot.Plot, index int, name string, xyers plotter.XYs) error {
	var ps []plot.Plotter

	s, err := plotter.NewScatter(xyers)
	if err != nil {
		return err
	}
	s.Color = plotutil.Color(index)
	s.Shape = plotutil.Shape(index)
	ps = append(ps, s)

	plt.Add(ps...)

	plt.Legend.Add(name)
	return nil
}
