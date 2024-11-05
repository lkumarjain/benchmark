package main

import (
	"flag"
	"fmt"

	"github.com/Knetic/govaluate"
	result "github.com/lkumarjain/benchmark-result"
)

func main() {
	parserFilePath := flag.String("ParserFilePath", "", "Parser file path")
	parserFileName := flag.String("ParserFileName", "results.out", "Parser file name")
	keyTemplate := flag.String("KeyTemplate", "{0}/{1}/{2}", "Key template where 0 - libraryName, 1 - scenarioName, 2 - functionName")
	legendTemplate := flag.String("LegendTemplate", "{1}/{2}", "Legend template where 0 - libraryName, 1 - scenarioName, 2 - functionName")
	optionsTemplate := flag.String("OptionsTemplate", "{0}", "Options template where 0 - libraryName, 1 - scenarioName, 2 - functionName")
	tableWidth := flag.Int("TableWidth", 2048, "Table width")
	tableLegendSpan := flag.Int("TableLegendSpan", 3, "Table legend span")
	barChartWidth := flag.Int("BarChartWidth", 4096, "Bar chart width")
	barChartHeight := flag.Int("BarChartHeight", 2160, "Bar chart height")
	filterExpression := flag.String("Filter", "", "Filter Expression")
	outputFileTemplate := flag.String("OutputFileTemplate", "%s", "Output file name prefix")

	flag.Parse()

	benchmarks := result.NewResult(*parserFilePath, *parserFileName, *keyTemplate, *legendTemplate, *optionsTemplate)
	benchmarks.TableWidth = *tableWidth
	benchmarks.TableLegendSpan = *tableLegendSpan
	benchmarks.BarChartWidth = *barChartWidth
	benchmarks.BarChartHeight = *barChartHeight
	benchmarks.OutputFileTemplate = *outputFileTemplate

	expression := *filterExpression
	if expression != "" {
		fmt.Println("Filter expression: ", expression)
		program, err := govaluate.NewEvaluableExpression(expression)
		if err != nil {
			panic(err)
		}

		benchmarks.Filter = program
	}

	err := benchmarks.Parse()
	if err != nil {
		panic(err)
	}
}
