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
	dataType        string
	minimum         []float64
	maximum         []float64
}

func newTable(width int, legendSpan int, options []string, legends []string, legendTemplate string, optionsTemplate string, dataType string) *table {
	headers := []string{"Options"}
	headers = append(headers, options...)

	return &table{
		width: width, legendSpan: legendSpan,
		headers: headers, data: make([][]string, len(legends)), legends: legends,
		legendTemplate: legendTemplate, optionsTemplate: optionsTemplate, dataType: dataType,
		minimum: make([]float64, len(options)), maximum: make([]float64, len(options)),
	}
}

func (t *table) addBenchmark(b benchmark) {
	legendKey := b.legendKey(t.legendTemplate)
	legendsIndex := slices.Index(t.legends, legendKey)

	optionsKey := b.optionsKey(t.optionsTemplate)
	optionsIndex := slices.Index(t.headers, optionsKey)

	if t.data[legendsIndex] == nil {
		t.data[legendsIndex] = make([]string, len(t.headers))
		t.data[legendsIndex][0] = legendKey
	}

	var value float64 = -1

	switch t.dataType {
	case TimeDataType:
		value = b.timePerOperation
		t.data[legendsIndex][optionsIndex] = duration(value)
	case MemoryDataType:
		value = b.memoryPerOperation
		t.data[legendsIndex][optionsIndex] = allocations(value)
	case AllocationsDataType:
		value = b.allocationsPerOperation
		t.data[legendsIndex][optionsIndex] = allocations(value)
	}

	if t.minimum[optionsIndex-1] == 0 || value <= t.minimum[optionsIndex-1] {
		t.minimum[optionsIndex-1] = value
	}

	if value >= t.maximum[optionsIndex-1] {
		t.maximum[optionsIndex-1] = value
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
		CellTextStyle: func(tc charts.TableCell) *charts.Style {
			index := tc.Column
			if tc.Row == 0 || index == 0 {
				return &tc.Style
			}

			min, max := t.minMax(index)

			if min != "" && tc.Text == min {
				return &charts.Style{
					FontColor: charts.Color{R: 33, G: 124, B: 50, A: 255},
					FontSize:  tc.Style.FontSize,
				}
			}

			if max != "" && tc.Text == max {
				return &charts.Style{
					FontColor: charts.Color{R: 179, G: 53, B: 20, A: 255},
					FontSize:  tc.Style.FontSize,
				}
			}

			return &tc.Style
		},
	}

	p, err := charts.TableOptionRender(options)
	if err != nil {
		panic(err)
	}

	return p.Bytes()
}

func (t *table) minMax(index int) (string, string) {
	switch t.dataType {
	case TimeDataType:
		return duration(t.minimum[index-1]), duration(t.maximum[index-1])
	case MemoryDataType:
		return allocations(t.minimum[index-1]), allocations(t.maximum[index-1])
	case AllocationsDataType:
		return allocations(t.minimum[index-1]), allocations(t.maximum[index-1])
	}

	return "", ""
}
