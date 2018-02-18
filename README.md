# Surrogate-Assisted Interactive GA Recommender 
[![Build Status](https://travis-ci.org/philipp-altmann/SAGRS.svg?branch=master)](https://travis-ci.org/philipp-altmann/SAGRS) [![Coverage Status](https://coveralls.io/repos/github/philipp-altmann/ContinuousBenchmarkOptimizer/badge.svg?branch=master)](https://coveralls.io/github/philipp-altmann/ContinuousBenchmarkOptimizer?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/philipp-altmann/ContinuousBenchmarkOptimizer)](https://goreportcard.com/report/github.com/philipp-altmann/ContinuousBenchmarkOptimizer)

## Run Benchmarks
```$ go run main.go <mode> <specific> ```

**mode**: benchmark  
**specific**: rates | suggestions | compare

### Evaluation Rate Testsuite
```$ go run main.go benchmark rates```

### Suggestions Testsuite
```$ go run main.go benchmark suggestions```

### Comparison Testsuite
comparing SAGRS to a Genetic Algorithm and a Random Search  
```$ go run main.go benchmark suggestions```
