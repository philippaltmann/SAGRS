package Utility

import (
	"fmt"
	"strings"

	"github.com/buger/goterm"
	a "github.com/philipp-altmann/SAGRS/Approximation"
	e "github.com/philipp-altmann/SAGRS/Environment"
	p "github.com/philipp-altmann/SAGRS/Population"
)

func writeConsole(e e.Environment, cycle, better, worse int, evaluationPool p.Population, Approximator a.Approximator) {
	//Print Stats
	goterm.MoveCursor(1, 1)
	betterString := goterm.Color(fmt.Sprintf("[%s]Better", getBetterLamp(better)), goterm.CYAN)
	goterm.Printf("%s[%4d]\n", makeProgressBar(cycle, e.Cycles, 50, goterm.GREEN), cycle)
	goterm.Printf("%s   Suggestion Success: %s\n", betterString, makeProgressBar(e.SuggestToEvaluation-worse, e.SuggestToEvaluation, 20, goterm.BLUE))
	goterm.Printf("Best in Pool at %6.6f       |Approximated at %6.6f\n", evaluationPool[0].Fitness, Approximator.Predict(evaluationPool[0].Value))
	goterm.Printf("             at Posititon (%6.6f / %6.6f)\n", evaluationPool[0].Value[0], evaluationPool[0].Value[1])
	//Print Approximator Stats

	goterm.Printf("Rejected %d of %d\n", worse, e.SuggestToEvaluation)
	goterm.Flush() // Call it every time at the end of rendering
}

func getBetterLamp(better int) string {
	if better > 0 {
		return "•"
	}
	return " "
}

func makeProgressBar(i, from, length int, color int) string {
	if from == 0 {
		from = 10
	}
	done := int(float64(i) / float64(from) * float64(length))
	return "[" + goterm.Color(strings.Repeat("=", done), color) + strings.Repeat("-", length-done) + "]"
}
