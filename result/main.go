package main

import (
	"os"
	"path/filepath"
	"slices"

	"github.com/lkumarjain/benchmark/result/parser"
	"github.com/lkumarjain/benchmark/result/plotter"
)

func main() {
	path := os.Args[1]
	records, applications, scenarios, err := parser.Parse(path)
	if err != nil {
		panic(err)
	}

	directory := filepath.Dir(path)

	allocationPerOperation := plotter.Plot{Time: false, Title: "allocationPerOperation", Directory: directory, FileName: "allocationPerOperation.png", Options: applications, LegendLabels: scenarios, Values: make([][]float64, len(scenarios))}
	timePerOperation := plotter.Plot{Time: true, Title: "timePerOperation", Directory: directory, FileName: "timePerOperation.png", Options: applications, LegendLabels: scenarios, Values: make([][]float64, len(scenarios))}

	for _, v := range records {
		index := slices.Index(scenarios, v.ScenarioName)
		allocations := allocationPerOperation.Values[index]
		if allocations == nil {
			allocations = make([]float64, len(applications))
			allocationPerOperation.Values[index] = allocations
		}

		allocations[slices.Index(applications, v.ApplicationName)] = v.AllocationPerOperation

		times := timePerOperation.Values[index]
		if times == nil {
			times = make([]float64, len(applications))
			timePerOperation.Values[index] = times
		}

		times[slices.Index(applications, v.ApplicationName)] = v.TimePerOperation
	}

	err = allocationPerOperation.Generate()
	if err != nil {
		panic(err)
	}

	err = timePerOperation.Generate()
	if err != nil {
		panic(err)
	}
}
