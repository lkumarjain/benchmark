package kafkaclient

import (
	"fmt"
	"strings"
)

var (
	bootstrapServers = "localhost:9092"
	userName         = "XXXXX"
	password         = "XXXXX"
	authenticator    = false
)

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

func topicName(prefix string) string {
	return fmt.Sprintf("test_broker_%s", strings.ToLower(prefix))
}
