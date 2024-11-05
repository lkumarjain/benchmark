package result

import (
	"github.com/wissance/stringFormatter"
)

type benchmark struct {
	libraryName             string
	scenarioName            string
	functionName            string
	iterations              int
	timePerOperation        float64
	timeUnit                string
	memoryPerOperation      float64
	memoryUnit              string
	allocationsPerOperation float64
	allocationsUnit         string
}

func (b benchmark) key(template string) string {
	return stringFormatter.Format(template, b.libraryName, b.scenarioName, b.functionName)
}

func (b benchmark) legendKey(template string) string {
	return stringFormatter.Format(template, b.libraryName, b.scenarioName, b.functionName)
}

func (b benchmark) optionsKey(template string) string {
	return stringFormatter.Format(template, b.libraryName, b.scenarioName, b.functionName)
}

func (b benchmark) Get(name string) (interface{}, error) {
	switch name {
	case "libraryName":
		return b.libraryName, nil
	case "scenarioName":
		return b.scenarioName, nil
	case "functionName":
		return b.functionName, nil
	case "iterations":
		return b.iterations, nil
	case "timePerOperation":
		return b.timePerOperation, nil
	case "timeUnit":
		return b.timeUnit, nil
	case "allocationPerOperation":
		return b.memoryPerOperation, nil
	case "allocationUnit":
		return b.memoryUnit, nil
	default:
		return nil, nil
	}
}
