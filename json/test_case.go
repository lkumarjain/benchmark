package json

import "fmt"

var tests = []struct {
	name    string
	payload string
	size    int
}{
	{name: "01KB", payload: payloadSmall, size: 15},
	{name: "05KB", payload: payloadMedium, size: 45},
	{name: "10KB", payload: payloadLarge, size: 85},
}

func testName(prefix string, function string) string {
	return fmt.Sprintf("%s/%s", prefix, function)
}
