package GeneticOptimizer

import (
	a "github.com/philipp-altmann/SAGRS/Approximation"
	e "github.com/philipp-altmann/SAGRS/Environment"
	i "github.com/philipp-altmann/SAGRS/Individual"
	o "github.com/philipp-altmann/SAGRS/Objective"
	p "github.com/philipp-altmann/SAGRS/Population"
	u "github.com/philipp-altmann/SAGRS/Utility"
)

//Optimize runs System for given Environment
func Optimize(e e.Environment) (bestIndividual i.Individual) {

	//Generate Approximator & Objective
	Objective := o.GetObjective(e.Objective)

	Approximator := a.GetApproximator(e.Objective)

	//Init, Evaluate & Sort Evaluation Pool randomly
	population := p.InitRandomPopulation(e.PopulationSize, e.Dimensions, Objective.Min(), Objective.Max())
	population.Evaluate(Objective.EvaluateFitness)

	//Setup Writer
	u.Setup(e)

	//Recommender Cycles
	for cycle := 0; cycle < e.Cycles; cycle++ {
		//Select
		population.Select(e.SelectionFactor)

		//Mutate
		population.Mutate(e.MutationFactor)

		//Recombine
		population.Recombine(e.RecombinationFactor, e.PopulationSize)

		//Fillup
		population.Fillup(e.PopulationSize, e.Dimensions, Objective.Min(), Objective.Max())

		//Evaluate using Approximator Function
		population.Evaluate(Objective.EvaluateFitness)

		//Write Status to Console & CSV
		u.WriteProgress(e, cycle, population, p.Population{}, Approximator)

	}
	return population[0]
}
