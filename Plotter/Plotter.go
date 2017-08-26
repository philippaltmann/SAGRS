package Plotter

import (
	"image/color"
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/philipp-altmann/ContinuousBenchmarkOptimizer/Generation"
)

func Plot(evaluated plotter.XYs, approximatedE plotter.XYs, approximatedA plotter.XYs) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Bohachevsky Optimization"
	p.X.Label.Text = "Cycle"
	p.Y.Label.Text = "Fitness"

	// Make a line plotter and set its style.
	le, err := plotter.NewLine(evaluated)
	if err != nil {
		panic(err)
	}
	le.LineStyle.Width = vg.Points(1)
	le.LineStyle.Color = color.RGBA{R: 255, A: 255}

	/*err = plotutil.AddLines(p,
		"Evaluated", evaluated)
	if err != nil {
		panic(err)
	}*/

	//TODO plot approxmation as offset

	la1, err := plotter.NewLine(approximatedE)
	if err != nil {
		panic(err)
	}
	la1.LineStyle.Width = vg.Points(1)
	la1.LineStyle.Color = color.RGBA{B: 255, A: 255}

	la2, err := plotter.NewLine(approximatedA)
	if err != nil {
		panic(err)
	}
	la2.LineStyle.Width = vg.Points(1)
	la2.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	la2.LineStyle.Color = color.RGBA{B: 255, A: 255}

	/*err = plotutil.AddLines(p,
		"Approximated(Appoximated)", approximatedA)
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLines(p,
		"Approximated(Evaluated)", approximatedE)
	if err != nil {
		panic(err)
	}*/

	p.Add(le, la1, la2)
	p.Legend.Add("Elaluated", le)
	p.Legend.Add("Approximated", la1, la2)

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "Plotter/Plots/points.svg"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func ConvertProgress(progress []float64) plotter.XYs {
	pts := make(plotter.XYs, len(progress))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = progress[i]
	}
	return pts
}

func Plot2D(gerneration Generation.Generation, idx int) {
	pointData := make(plotter.XYs, len(gerneration))
	for i := range pointData {
		pointData[i].X = gerneration[i].Value[0]
		pointData[i].Y = gerneration[i].Value[1]
	}

	// Create a new plot, set its title and
	// axis labels.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(pointData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	/*
		// Make a line plotter and set its style.
		l, err := plotter.NewLine(lineData)
		if err != nil {
			panic(err)
		}
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
		l.LineStyle.Color = color.RGBA{B: 255, A: 255}

		// Make a line plotter with points and set its style.
		lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
		if err != nil {
			panic(err)
		}
		lpLine.Color = color.RGBA{G: 255, A: 255}
		lpPoints.Shape = draw.PyramidGlyph{}
		lpPoints.Color = color.RGBA{R: 255, A: 255}
	*/
	// Add the plotters to the plot, with a legend
	// entry for each
	p.Add(s)
	//p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	//p.Legend.Add("line", l)
	//p.Legend.Add("line points", lpLine, lpPoints)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "Plotter/ValueChange/values"+strconv.Itoa(idx)+".png"); err != nil {
		panic(err)
	}
}
