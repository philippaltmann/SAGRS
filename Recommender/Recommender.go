package Recommender

import (
	a "github.com/philipp-altmann/SAGRS/Approximation"
	e "github.com/philipp-altmann/SAGRS/Environment"
	i "github.com/philipp-altmann/SAGRS/Individual"
	o "github.com/philipp-altmann/SAGRS/Objective"
	p "github.com/philipp-altmann/SAGRS/Population"
	u "github.com/philipp-altmann/SAGRS/Utility"
)

//Run Recommender System for given Environment
func Run(e e.Environment) (bestIndividual i.Individual, cycle int) {

	//Generate Approximator & Objective
	Objective := o.GetObjective(e.Objective)
	Approximator := a.GetApproximator(e.Approximator)

	//Init, Evaluate & Sort Evaluation Pool randomly
	evaluationPool := p.InitRandomPopulation(e.EvaluationPoolSize, e.Dimensions, Objective.Min(), Objective.Max())
	evaluationPool.Evaluate(Objective.EvaluateFitness)

	//Update Approximator
	Approximator.Update(evaluationPool)

	//Setup Writer
	u.Setup(e)

	//Init, Approximate & Sort Population
	population := p.InitRandomPopulation(e.PopulationSize, e.Dimensions, Objective.Min(), Objective.Max())
	population.Evaluate(Approximator.Predict)

	//Recommender Cycles
	for cycle := 0; cycle < e.Cycles; cycle++ {
		if e.ResetPool {
			//Init, Approximate & Sort Population
			population = p.InitRandomPopulation(e.PopulationSize-1, e.Dimensions, Objective.Min(), Objective.Max())
			population.Evaluate(Approximator.Predict)
		}

		//Optimize Population for Suggestion using Approximator
		for generation := 0; generation < e.EvaluationRate; generation++ {

			//Select
			population.Select(e.SelectionFactor)

			//Mutate
			population.Mutate(e.MutationFactor)

			//Recombine
			population.Recombine(e.RecombinationFactor, e.PopulationSize)

			//Fillup
			population.Fillup(e.PopulationSize, e.Dimensions, Objective.Min(), Objective.Max())

			//Evaluate using Approximator Function
			population.Evaluate(Approximator.Predict)
		}

		//Suggest
		var suggestions []i.Individual
		for s := 0; s < e.SuggestToEvaluation; s++ {
			suggestions = append(suggestions, i.DuplicateIndividual(population[s]))
		}

		//Write Status to Console & CSV
		u.WriteProgress(e, cycle, evaluationPool, suggestions, Approximator)

		//Recomment Suggestions
		evaluationPool = append(evaluationPool, suggestions...)
		evaluationPool.Evaluate(Objective.EvaluateFitness)

		//Cut worst to keep Evaluation Pool Size Constant
		evaluationPool = evaluationPool[:e.EvaluationPoolSize]

		//Update Approximator
		Approximator.Update(evaluationPool)

	}
	convergenceCycle := u.GetConvergenceCycle()
	return evaluationPool[0], convergenceCycle //evaluationPool
}
