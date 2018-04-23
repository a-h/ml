package projection

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/a-h/raster"
	"golang.org/x/image/colornames"
)

// A Projection can write out a graph.
type Projection struct {
	// The size of the image to draw
	Width, Height int
	// The minimum x and y values.
	MinimumValue float64
	// The maximum x and y values.
	MaximumValue float64
	// The number of cells to render.
	Cells int
	// The function used to produce the z value.
	Function func(x, y float64) float64
	// The angle of the projection
	Angle    float64
	cosAngle float64
	sinAngle float64
	// The Z scale of the projection
	Scale           float64
	BackgroundColor color.RGBA
	LineColor       color.RGBA
	FillColor       color.RGBA
}

func New(min, max float64, function func(x, y float64) float64, width, height int, angle float64) Projection {
	rads := math.Pi / (float64(180) / angle)
	return Projection{
		Width:           width,
		Height:          height,
		MinimumValue:    min,
		MaximumValue:    max,
		Function:        function,
		Cells:           100,
		Angle:           rads,
		cosAngle:        math.Cos(rads),
		sinAngle:        math.Sin(rads),
		Scale:           float64(height) * 0.4,
		BackgroundColor: color.RGBA{},
		LineColor:       colornames.Blue,
		FillColor:       colornames.Lightblue,
	}
}

func (p Projection) Draw(img draw.Image) {
	for i := 0; i < p.Cells; i++ {
		for j := 0; j < p.Cells; j++ {
			ax, ay, ok := corner(p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i+1, j, p.Cells, p.Function)
			if !ok {
				continue
			}
			bx, by, ok := corner(p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i, j, p.Cells, p.Function)
			if !ok {
				continue
			}
			cx, cy, ok := corner(p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i, j+1, p.Cells, p.Function)
			if !ok {
				continue
			}
			dx, dy, ok := corner(p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i+1, j+1, p.Cells, p.Function)
			if !ok {
				continue
			}

			a := image.Point{int(ax), int(ay)}
			b := image.Point{int(bx), int(by)}
			c := image.Point{int(cx), int(cy)}
			d := image.Point{int(dx), int(dy)}

			pg := raster.NewFilledPolygon(p.LineColor, p.FillColor, a, b, c, d)
			pg.Draw(img)
		}
	}
}

func corner(width, height int, cosAngle, sinAngle float64, zscale float64, i, j int, cells int, f func(float64, float64) float64) (float64, float64, bool) {
	xyrange := 30.0
	xyscale := float64(width) / 2.0 / xyrange

	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cosAngle*xyscale //- z*zscale
	sy := float64(height)/2 + (x+y)*sinAngle*xyscale - z*zscale

	if math.IsNaN(sx) || math.IsNaN(sy) {
		return 0, 0, false
	}

	return sx, sy, true
}

func drawBackground(img *image.RGBA, c color.RGBA) {
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, c)
		}
	}
}
