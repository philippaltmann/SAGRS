package Environment

import (
	"encoding/json"
	"io/ioutil"
)

//Environment settings
type Environment struct {
	Cycles              int     // = 5000
	EvaluationRate      int     //= 100
	SuggestToEvaluation int     // = 1
	EvaluationPoolSize  int     // = 10
	PopulationSize      int     // = 100
	Dimensions          int     //= 2
	SelectionFactor     float64 //= 0.6
	MutationFactor      float64 //= 0.3
	RecombinationFactor float64 //= 0.4
	ProgressFileName    string
	WriteProgress       bool
	Verbose             bool
	ResetPool           bool
	Approximator        string
	Objective           string
}

//Dump Formats and writes current Environment to json File
func (e Environment) Dump(path string) {
	env, _ := json.Marshal(e)
	ioutil.WriteFile(path, env, 0644)
}
