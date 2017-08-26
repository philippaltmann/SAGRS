package Plotter

import (
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
