package parser

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Benchmark struct {
	ApplicationName        string
	ScenarioName           string
	NumberOfCPU            string
	Iterations             int
	TimePerOperation       float64
	TimeUnit               string
	AllocationPerOperation float64
	AllocationUnit         string
}

func (b *Benchmark) Parse(line string) error {
	if strings.HasPrefix(line, "Benchmark") {
		fields := strings.Fields(line[9:])

		names := strings.Split(fields[0], "/")
		b.ApplicationName = names[0]
		b.ScenarioName = strings.ReplaceAll(names[1][:len(names[1])-2], "_", " ")
		b.NumberOfCPU = names[1][len(names[1])-1:]
		b.Iterations, _ = strconv.Atoi(fields[1])
		b.TimePerOperation, _ = strconv.ParseFloat(fields[2], 64)
		b.TimeUnit = fields[3]
		b.AllocationPerOperation, _ = strconv.ParseFloat(fields[4], 64)
		b.AllocationUnit = fields[5]

		return nil
	}

	return fmt.Errorf("invalid line: %s", line)
}

func Parse(path string) ([]Benchmark, []string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var benchmarks = make(map[string]Benchmark)

	for scanner.Scan() {
		line := scanner.Text()
		b := &Benchmark{}

		err := b.Parse(line)
		if err != nil {
			continue
		}

		key := b.ApplicationName + "/" + b.ScenarioName
		benchmark, ok := benchmarks[key]
		if !ok {
			benchmark = *b
		} else {
			benchmark.Iterations = (benchmark.Iterations + b.Iterations) / 2
			benchmark.TimePerOperation = (benchmark.TimePerOperation + b.TimePerOperation) / 2
			benchmark.AllocationPerOperation = (benchmark.AllocationPerOperation + b.AllocationPerOperation) / 2
		}

		benchmarks[key] = benchmark

		key = b.ApplicationName
		benchmark, ok = benchmarks[key]
		if !ok {
			benchmark = *b
			benchmark.ScenarioName = "All"
		} else {
			benchmark.Iterations = (benchmark.Iterations + b.Iterations) / 2
			benchmark.TimePerOperation = (benchmark.TimePerOperation + b.TimePerOperation) / 2
			benchmark.AllocationPerOperation = (benchmark.AllocationPerOperation + b.AllocationPerOperation) / 2
		}

		benchmarks[key] = benchmark
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, errors.New("failed to scan file")
	}

	var records []Benchmark
	var applications []string
	var scenarios []string

	for _, v := range benchmarks {
		v.TimePerOperation = math.Round(v.TimePerOperation*100) / 100
		v.AllocationPerOperation = math.Round(v.AllocationPerOperation*100) / 100

		records = append(records, v)

		if !slices.Contains(applications, v.ApplicationName) {
			applications = append(applications, v.ApplicationName)
		}

		if !slices.Contains(scenarios, v.ScenarioName) {
			scenarios = append(scenarios, v.ScenarioName)
		}
	}

	slices.Sort(applications)
	slices.Sort(scenarios)

	return records, applications, scenarios, nil
}
