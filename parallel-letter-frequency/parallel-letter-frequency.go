// Package letter provides a function to run a frequency analysis function in parallel.
package letter

import (
	"flag"
	"sync"
)

type count struct {
	k rune
	v int
}

var (
	depth int
)

func init() {
	depthFlag := flag.Int("depth", 16, "buffer depth") // based on some quick ad-hoc benchmarking (see code below)
	flag.Parse()
	depth = *depthFlag
}

// ConcurrentFrequency calculates frequency of letters in a slice of strings, concurrently
func ConcurrentFrequency(sources []string) FreqMap {
	var con = make(FreqMap)
	var ch = make(chan count, depth)
	var wg sync.WaitGroup

	// for each source string
	for _, source := range sources {
		wg.Add(1)
		go func(source string) {
			defer wg.Done()
			freq := Frequency(source)
			for k, v := range freq {
				ch <- count{k, v}
			}
		}(source)
	}

	// close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// gather results until channel is closed
	for c := range ch {
		con[c.k] += c.v
	}

	return con
}

/* to benchmark, add to the test suite:

func BenchmarkConcurrentFrequency(b *testing.B) {
	b.StopTimer()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ConcurrentFrequency([]string{euro, dutch, us})
		}

		b.StopTimer()
}

*/
