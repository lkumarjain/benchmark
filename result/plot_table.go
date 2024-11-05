package result

import (
	"slices"

	"github.com/vicanso/go-charts/v2"
)

type table struct {
	width           int
	legends         []string
	headers         []string
	data            [][]string
	legendSpan      int
	legendTemplate  string
	optionsTemplate string
}

func newTable(width int, legendSpan int, options []string, legends []string, legendTemplate string, optionsTemplate string) *table {
	headers := []string{"Options"}
	headers = append(headers, options...)

	return &table{
		width: width, legendSpan: legendSpan,
		headers: headers, data: make([][]string, len(legends)), legends: legends,
		legendTemplate: legendTemplate, optionsTemplate: optionsTemplate,
	}
}

func (t *table) addBenchmark(b benchmark, dataType string) {
	legendKey := b.legendKey(t.legendTemplate)
	legendsIndex := slices.Index(t.legends, legendKey)

	optionsKey := b.optionsKey(t.optionsTemplate)
	optionsIndex := slices.Index(t.headers, optionsKey)

	if t.data[legendsIndex] == nil {
		t.data[legendsIndex] = make([]string, len(t.headers))
		t.data[legendsIndex][0] = legendKey
	}

	switch dataType {
	case TimeDataType:
		t.data[legendsIndex][optionsIndex] = duration(b.timePerOperation)
	case MemoryDataType:
		t.data[legendsIndex][optionsIndex] = allocations(b.memoryPerOperation)
	case AllocationsDataType:
		t.data[legendsIndex][optionsIndex] = allocations(b.allocationsPerOperation)
	}
}

func (t *table) plot() ([]byte, error) {
	spans := make([]int, len(t.headers))
	spans[0] = t.legendSpan
	for i := range t.headers {
		if i == 0 {
			continue
		}

		spans[i] = 1
	}

	options := charts.TableChartOption{
		Header:                t.headers,
		Data:                  t.data,
		Width:                 t.width,
		FontSize:              30,
		HeaderBackgroundColor: charts.Color{R: 16, G: 22, B: 30, A: 255},
		HeaderFontColor:       charts.Color{R: 255, G: 255, B: 255, A: 255},
		Spans:                 spans,
	}

	p, err := charts.TableOptionRender(options)
	if err != nil {
		panic(err)
	}

	return p.Bytes()
}
