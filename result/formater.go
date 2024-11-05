package result

import (
	"fmt"
	"time"
)

const (
	// TimeDataType time data type
	TimeDataType = "time"
	// MemoryDataType memory data type
	MemoryDataType = "memory"
	// AllocationsDataType allocations data type
	AllocationsDataType = "allocations"
)

func duration(f float64) string {
	return time.Duration(int64(f)).String()
}

func allocations(value float64) string {
	if value < 1024 {
		return fmt.Sprintf("%.0f B", value)
	}

	if value < 1048576 {
		return fmt.Sprintf("%.0f KB", value/1024)
	}

	return fmt.Sprintf("%.0f MB", value/1048576)
}
