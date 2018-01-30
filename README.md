# Surrogate-Assisted Interactive GA Recommender [![Build Status](https://travis-ci.org/philipp-altmann/ContinuousBenchmarkOptimizer.svg?branch=master)](https://travis-ci.org/philipp-altmann/ContinuousBenchmarkOptimizer) [![Coverage Status](https://coveralls.io/repos/github/philipp-altmann/ContinuousBenchmarkOptimizer/badge.svg?branch=master)](https://coveralls.io/github/philipp-altmann/ContinuousBenchmarkOptimizer?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/philipp-altmann/ContinuousBenchmarkOptimizer)](https://goreportcard.com/report/github.com/philipp-altmann/ContinuousBenchmarkOptimizer)

## Run Benchmarks
```$ go run main.go <mode> <specific> <approximation model> <objective>```

**mode**: benchmark  
**specific**: rates | convergence | suggestions  
**approximation model**: RBF | LSM  
**objective**: Bohachevsky | Ackley | Schwefel

### Evaluation Rate Testsuite
```$ go run main.go benchmark rates```
