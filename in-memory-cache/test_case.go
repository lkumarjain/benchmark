package inmemorycache

import "fmt"

var tests = []struct {
	name           string
	valueGenerator func(int) string
}{
	{name: "1KB", valueGenerator: generateSmall},
	{name: "5KB", valueGenerator: generateMedium},
	{name: "10KB", valueGenerator: generateLarge},
}

func generateKey(prefix string, index int) string {
	return fmt.Sprintf("%s:%d", prefix, index)
}

func generateSmall(index int) string {
	return fmt.Sprintf("%s:%d", payloadSmall, index)
}

func generateMedium(index int) string {
	return fmt.Sprintf("%s:%d", payloadMedium, index)
}

func generateLarge(index int) string {
	return fmt.Sprintf("%s:%d", payloadLarge, index)
}

func testName(prefix string, function string) string {
	return fmt.Sprintf("%s - %s", prefix, function)
}
