package result

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

type Result struct {
	ParserFilePath     string
	ParserFileName     string
	KeyTemplate        string
	LegendTemplate     string
	OptionsTemplate    string
	Filter             *govaluate.EvaluableExpression
	TableWidth         int
	TableLegendSpan    int
	BarChartWidth      int
	BarChartHeight     int
	OutputFileTemplate string
	benchmarks         []benchmark
	legends            []string
	options            []string
}

func NewResult(parserFilePath string, parserFileName string, keyTemplate string, legendTemplate string, optionsTemplate string) *Result {
	return &Result{
		ParserFilePath: parserFilePath, ParserFileName: parserFileName,
		KeyTemplate: keyTemplate, LegendTemplate: legendTemplate, OptionsTemplate: optionsTemplate,
		TableWidth: 2048, TableLegendSpan: 3, BarChartWidth: 4096, BarChartHeight: 2160,
	}
}

func (r *Result) Parse() error {
	file, err := os.Open(filepath.Join(r.ParserFilePath, r.ParserFileName))
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	r.benchmarks = make([]benchmark, 0)
	r.legends = make([]string, 0)
	r.options = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}

		r.parseLine(line)
	}

	slices.SortFunc(r.benchmarks, func(a, b benchmark) int {
		if a.libraryName < b.libraryName {
			return -1
		}

		if a.libraryName > b.libraryName {
			return 1
		}

		return 0
	})

	slices.Sort(r.legends)
	slices.Sort(r.options)

	for _, result := range r.benchmarks {
		fmt.Println(result)
	}

	fmt.Println(r.legends)

	return r.Plot()
}

func (r *Result) parseLine(line string) {
	b := benchmark{}

	fields := strings.Fields(line[9:])
	names := strings.Split(fields[0], "/")
	b.libraryName = names[0]
	b.scenarioName = names[1]
	b.functionName = strings.ReplaceAll(names[2][:len(names[2])-2], "_", " ")
	b.iterations, _ = strconv.Atoi(fields[1])
	b.timePerOperation, _ = strconv.ParseFloat(fields[2], 64)
	b.timeUnit = fields[3]
	b.memoryPerOperation, _ = strconv.ParseFloat(fields[4], 64)
	b.memoryUnit = fields[5]
	b.allocationsPerOperation, _ = strconv.ParseFloat(fields[6], 64)
	b.allocationsUnit = fields[7]

	filter := false
	if r.Filter != nil {
		out, err := r.Filter.Eval(b)
		if err == nil && out.(bool) {
			filter = true
		}
	}

	if !filter {
		r.addBenchmark(b)
	}
}

func (r *Result) addBenchmark(b benchmark) {
	index := slices.IndexFunc(r.benchmarks, func(value benchmark) bool {
		return b.key(r.KeyTemplate) == value.key(r.KeyTemplate)
	})

	if index < 0 {
		r.benchmarks = append(r.benchmarks, b)
		index = len(r.benchmarks) - 1
	}

	benchmark := r.benchmarks[index]
	benchmark.iterations = (benchmark.iterations + b.iterations) / 2
	benchmark.timePerOperation = (benchmark.timePerOperation + b.timePerOperation) / 2
	benchmark.memoryPerOperation = (benchmark.memoryPerOperation + b.memoryPerOperation) / 2

	benchmark.timePerOperation = math.Round(benchmark.timePerOperation*100) / 100
	benchmark.memoryPerOperation = math.Round(benchmark.memoryPerOperation*100) / 100

	r.benchmarks[index] = benchmark

	legendKey := b.legendKey(r.LegendTemplate)

	legendsIndex := slices.Index(r.legends, legendKey)
	if legendsIndex < 0 {
		r.legends = append(r.legends, legendKey)
	}

	optionsKey := b.optionsKey(r.OptionsTemplate)

	optionsIndex := slices.Index(r.options, optionsKey)
	if optionsIndex < 0 {
		r.options = append(r.options, optionsKey)
	}
}

func (r *Result) Plot() error {
	timeTable := newTable(r.TableWidth, r.TableLegendSpan, r.legends, r.options, r.OptionsTemplate, r.LegendTemplate, TimeDataType)
	memoryTable := newTable(r.TableWidth, r.TableLegendSpan, r.legends, r.options, r.OptionsTemplate, r.LegendTemplate, MemoryDataType)
	allocationsTable := newTable(r.TableWidth, r.TableLegendSpan, r.legends, r.options, r.OptionsTemplate, r.LegendTemplate, AllocationsDataType)

	timeBar := newBar(r.BarChartWidth, r.BarChartHeight, r.options, r.legends, TimeDataType, r.LegendTemplate, r.OptionsTemplate)
	timeBar.title = "Time/Operation"

	memoryBar := newBar(r.BarChartWidth, r.BarChartHeight, r.options, r.legends, MemoryDataType, r.LegendTemplate, r.OptionsTemplate)
	memoryBar.title = "Memory/Operation"

	allocationsBar := newBar(r.BarChartWidth, r.BarChartHeight, r.options, r.legends, AllocationsDataType, r.LegendTemplate, r.OptionsTemplate)
	allocationsBar.title = "Allocations/Operation"

	for _, b := range r.benchmarks {
		timeTable.addBenchmark(b)
		memoryTable.addBenchmark(b)
		allocationsTable.addBenchmark(b)

		timeBar.addBenchmark(b)
		memoryBar.addBenchmark(b)
		allocationsBar.addBenchmark(b)
	}

	err := r.writeTables(timeTable, memoryTable, allocationsTable)
	if err != nil {
		return err
	}

	err = r.writeBar(timeBar, memoryBar, allocationsBar)

	return err

}

func (r *Result) writeTables(timeTable *table, memoryTable *table, allocationsTable *table) error {
	buf, err := timeTable.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "time_table.png"))
	if err != nil {
		return err
	}

	buf, err = memoryTable.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "memory_table.png"))
	if err != nil {
		return err
	}

	buf, err = allocationsTable.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "allocations_table.png"))

	return err
}

func (r *Result) writeBar(timeBar *bar, memoryBar *bar, allocationsBar *bar) error {
	buf, err := timeBar.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "time_bar.png"))
	if err != nil {
		return err
	}

	buf, err = memoryBar.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "memory_bar.png"))
	if err != nil {
		return err
	}

	buf, err = allocationsBar.plot()
	if err != nil {
		return err
	}

	err = r.writeChart(buf, fmt.Sprintf(r.OutputFileTemplate, "allocations_bar.png"))

	return err
}

func (r *Result) writeChart(buf []byte, fileName string) error {
	file := filepath.Join(r.ParserFilePath, fileName)

	err := os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}

	return nil
}
