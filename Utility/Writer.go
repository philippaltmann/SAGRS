package Utility

import (
	"strconv"

	a "github.com/philipp-altmann/SAGRS/Approximation"
	e "github.com/philipp-altmann/SAGRS/Environment"
	p "github.com/philipp-altmann/SAGRS/Population"
)

//TODO write tests

//WriteCSV writes current state to given CSV File
func writeCSV(e e.Environment, cycle, worse int, bestFitness float64, suggestions p.Population, approximator a.Approximator) {

	//Write Progress to csv
	line := make([]string, 4)
	//Cycle
	line[0] = strconv.Itoa(cycle)

	//Best in Pool
	line[1] = formatFloat(bestFitness)

	//Approximation Success
	line[2] = formatFloat(float64(e.SuggestToEvaluation-worse) / float64(e.SuggestToEvaluation))

	//Approximation Error
	line[3] = "" //formatFloat(approximator.Predict(suggestions[0].Value) - objective.EvaluateFitness(suggestions[0].Value))

	//Write & Close File
	writer.Write(line)
	writer.Flush()
}

func formatFloat(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}
