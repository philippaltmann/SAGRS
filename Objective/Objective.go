package Objective

import l "github.com/philipp-altmann/SAGRS/Objective/Linear"
import a "github.com/philipp-altmann/SAGRS/Objective/Ackley"
import b "github.com/philipp-altmann/SAGRS/Objective/Bohachevsky"
import s "github.com/philipp-altmann/SAGRS/Objective/Schwefel"

//Objective Interface for universal use of Objectives
type Objective interface {
	EvaluateFitness([]float64) float64
	Min() float64
	Max() float64
}

//GetObjective Generates Objective By Name
//Defaults to linear Objective used For Testing
func GetObjective(name string) Objective {
	switch name {
	case "Ackley":
		return a.Ackley{}
	case "Schwefel":
		return s.Schwefel{}
	case "Bohachevsky":
		return b.Bohachevsky{}
	default:
		return l.Linear{}
	}
}
