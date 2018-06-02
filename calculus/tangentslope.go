package calculus

// TangentSlope returns the slope of the line at point x by measuring the slope from x - d to x.
func TangentSlope(x, d float64, f func(x float64) (y float64)) Line {
	xi, xii := x-d, x
	yi, yii := f(xi), f(xii)

	rise, run := yii-yi, xii-xi
	m := rise / run

	// When x is equal to zero, we've hit the y-intercept.
	// Lets calculate b first.
	// yi = m(xi) + b
	// yi - m(xi) = b
	b := yi - (m * xi)

	return Line{
		M: m,
		B: b,
	}
}

// Line presents a slope intercept formula: y = mx + b
type Line struct {
	// M is the slope of the line.
	M float64
	// B is the y-intercept.
	B float64
}

// Y calculates the Y value.
func (l Line) Y(x float64) float64 {
	return (l.M * x) + l.B
}

// X calculates the X value.
func (l Line) X(y float64) float64 {
	// y = mx + b
	// y - b = mx
	// (y - b) / m = x
	return (y - l.B) / l.M
}
