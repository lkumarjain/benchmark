package plotter

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vicanso/go-charts/v2"
)

var (
	markLineOption  = charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage)
	markPointOption = charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax, charts.SeriesMarkDataTypeMin)
	imageType       = "png"
	theme           = "grafana"
	width           = 4096
	height          = 2160
	fontSize        = float64(30)
	padding         = charts.Box{Top: 100, Bottom: 100, Left: 100, Right: 300}
)

type Plot struct {
	Title        string
	Values       [][]float64
	Options      []string
	LegendLabels []string
	Directory    string
	FileName     string
	Time         bool
}

func (p *Plot) writeFile(buf []byte) error {
	err := os.MkdirAll(p.Directory, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(p.Directory, p.FileName)

	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (p *Plot) valueFormatter(f float64) string {
	if p.Time {
		return time.Duration(int64(f)).String()
	}

	return allocationUnitValue(f)
}

func (p *Plot) options(opt *charts.ChartOption) {
	if len(opt.SeriesList) > 0 {
		opt.SeriesList[1].MarkPoint = charts.NewMarkPoint(charts.SeriesMarkDataTypeMax, charts.SeriesMarkDataTypeMin)
		opt.SeriesList[1].MarkLine = charts.NewMarkLine(charts.SeriesMarkDataTypeAverage)
		opt.SeriesList[1].MarkLine.Width = 10
	}

	opt.Type = imageType
	opt.Theme = theme
	opt.Width = width
	opt.Height = height
	opt.Padding = padding
	opt.XAxis.FontSize = fontSize
	opt.Legend.FontSize = fontSize
	opt.BarMargin = 5
	opt.Title = charts.TitleOption{
		Text:            p.Title,
		Subtext:         "Benchmark Created by github.com/lkumarjain/benchmark",
		Left:            charts.PositionRight,
		Top:             charts.PositionTop,
		FontSize:        50,
		SubtextFontSize: 25,
	}

	opt.YAxisOptions = []charts.YAxisOption{
		{
			FontSize:      fontSize,
			Show:          charts.TrueFlag(),
			SplitLineShow: charts.TrueFlag(),
		},
	}

	opt.ValueFormatter = p.valueFormatter
}

func (p *Plot) Generate() error {
	xAxisOptions := charts.XAxisDataOptionFunc(p.Options)
	legendLabelsOption := charts.LegendLabelsOptionFunc(p.LegendLabels, charts.PositionLeft)

	chart, err := charts.BarRender(p.Values, xAxisOptions, legendLabelsOption, markLineOption, markPointOption, p.options)
	if err != nil {
		return err
	}

	buf, err := chart.Bytes()
	if err != nil {
		return err
	}

	err = p.writeFile(buf)
	if err != nil {
		return err
	}

	return nil
}

func allocationUnitValue(value float64) string {
	if value < 1024 {
		return fmt.Sprintf("%.0f B", value)
	}

	if value < 1048576 {
		return fmt.Sprintf("%.0f KB", value/1024)
	}

	return fmt.Sprintf("%.0f MB", value/1048576)
}
