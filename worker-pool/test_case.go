package workerpool

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"
)

var tests = []struct {
	name     string
	title    string
	executor func(ctx context.Context, id string) (string, error)
}{
	{
		name: "Sleep=001ms", title: "Sleep for 1 microsecond",
		executor: func(ctx context.Context, id string) (string, error) {
			time.Sleep(time.Microsecond)
			return "", nil
		},
	},
	{
		name: "Sleep=010ms", title: "Sleep for 10 microsecond",
		executor: func(ctx context.Context, id string) (string, error) {
			time.Sleep(10 * time.Microsecond)
			return id, nil
		},
	},
	{
		name: "SHA256=1kB", title: "SHA256 hash over 1kB",
		executor: func(ctx context.Context, id string) (string, error) {
			return fmt.Sprintf("%x", sha256.Sum256(oneKiloByte)), nil
		},
	},
	{
		name: "SHA256=8kB", title: "SHA256 hash over 8kB",
		executor: func(ctx context.Context, id string) (string, error) {
			return fmt.Sprintf("%x", sha256.Sum256(eightKiloByte)), nil
		},
	},
}
