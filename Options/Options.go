package Options

type Options struct {
	Cycles              int     // = 5000
	EvaluationRate      int     //= 100
	EvaluationPoolSize  int     // = 10
	PopulationSize      int     // = 100
	Dimensions          int     //= 2
	SelectionFactor     float64 //= 0.6
	MutationFactor      float64 //= 0.3
	RecombinationFactor float64 //= 0.4

}
