package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

const TEST_RUNS = 10

func execute(exe string) time.Duration {
	start := time.Now()
	cmd := exec.Command(exe)
	_ = cmd.Run()
	end := time.Since(start)
	return end
}

func averageDuration(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	var total time.Duration
	for _, duration := range durations {
		total += duration
	}

	return time.Duration(total.Nanoseconds() / int64(len(durations)))
}

func runTest(command, test string, runs int) {
	results := make([]time.Duration, 10)

	for i := 0; i < runs; i++ {
		result := execute(command)
		results = append(results, result)
	}

	average := averageDuration(results)

	fmt.Println(fmt.Sprintf("%s:", test), average)
}

func main() {
	testRunsFlag := flag.Int("runs", TEST_RUNS, "The amount of tests to run")
	flag.Parse()

	testRuns := *testRunsFlag

	testData, err := os.ReadFile("tests.json")
	if err != nil {
		log.Fatal(err)
	}

	var tests map[string]string
	err = json.Unmarshal(testData, &tests)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for k, v := range tests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runTest(v, k, testRuns)
		}()
	}

	wg.Wait()
}
