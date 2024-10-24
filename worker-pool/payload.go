package workerpool

import (
	"strings"
	"sync"
)

var (
	wg            sync.WaitGroup
	oneKiloByte   = []byte(strings.Repeat("a", 1024))
	eightKiloByte = []byte(strings.Repeat("a", 8192))
	concurrency   = []int{1, 10, 50, 100}
	maxPoolSize   = 100
)
