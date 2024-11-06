package result

import (
	"slices"

	"github.com/vicanso/go-charts/v2"
)

var (
	fontSize  = float64(30)
	imageType = "png"
	theme     = "grafana"
)

type bar struct {
	title           string
	data            [][]float64
	options         []string
	legends         []string
	width           int
	height          int
	dataType        string
	legendTemplate  string
	optionsTemplate string
}

func newBar(width int, height int, options []string, legends []string, dataType string, legendTemplate string, optionsTemplate string) *bar {
	return &bar{
		width: width, height: height,
		options: options, data: make([][]float64, len(legends)), legends: legends,
		dataType: dataType, legendTemplate: legendTemplate, optionsTemplate: optionsTemplate,
	}
}

func (ba *bar) addBenchmark(bm benchmark) {
	legendKey := bm.legendKey(ba.legendTemplate)
	legendsIndex := slices.Index(ba.legends, legendKey)

	optionsKey := bm.optionsKey(ba.optionsTemplate)
	optionsIndex := slices.Index(ba.options, optionsKey)

	if ba.data[legendsIndex] == nil {
		ba.data[legendsIndex] = make([]float64, len(ba.options))
	}

	switch ba.dataType {
	case TimeDataType:
		ba.data[legendsIndex][optionsIndex] = bm.timePerOperation
	case MemoryDataType:
		ba.data[legendsIndex][optionsIndex] = bm.memoryPerOperation
	case AllocationsDataType:
		ba.data[legendsIndex][optionsIndex] = bm.allocationsPerOperation
	}
}

func (ba *bar) plot() ([]byte, error) {
	xAxisOptions := charts.XAxisDataOptionFunc(ba.options)
	legendLabelsOption := charts.LegendLabelsOptionFunc(ba.legends)

	chart, err := charts.BarRender(ba.data, xAxisOptions, legendLabelsOption, ba.optionFunc)
	if err != nil {
		return nil, err
	}

	buf, err := chart.Bytes()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (ba *bar) valueFormatter(f float64) string {
	if ba.dataType == TimeDataType {
		return duration(f)
	}

	return allocations(f)
}

func (ba *bar) optionFunc(opt *charts.ChartOption) {
	for i := 0; i < len(opt.SeriesList); i++ {
		opt.SeriesList[i].MarkPoint = charts.NewMarkPoint(charts.SeriesMarkDataTypeMax, charts.SeriesMarkDataTypeMin)
		opt.SeriesList[i].MarkPoint.SymbolSize = 50
	}

	chartPadding := charts.Box{Top: 100, Bottom: 100, Left: 100, Right: 100}
	if len(ba.options) > 10 {
		chartPadding = charts.Box{Top: 100, Bottom: 500, Left: 100, Right: 100}
		opt.XAxis.TextRotation = 90
		opt.XAxis.LabelOffset = charts.Box{Top: -200, Bottom: 0, Left: 20, Right: 0}
	}

	opt.Type = imageType
	opt.Theme = theme
	opt.Width = ba.width
	opt.Height = ba.height
	opt.Padding = chartPadding

	opt.XAxis.FontSize = fontSize

	opt.Legend.Padding = charts.Box{Top: 200, Bottom: 200, Left: 300, Right: 100}
	opt.Legend.FontSize = fontSize

	opt.BarMargin = 10

	opt.Title = charts.TitleOption{
		Text:            ba.title,
		Subtext:         "Benchmark created by github.com/lkumarjain/benchmark",
		Left:            charts.PositionCenter,
		FontSize:        50,
		SubtextFontSize: 30,
	}

	opt.YAxisOptions = []charts.YAxisOption{
		{
			FontSize:      fontSize,
			Show:          charts.TrueFlag(),
			SplitLineShow: charts.TrueFlag(),
		},
	}

	opt.ValueFormatter = ba.valueFormatter
}
