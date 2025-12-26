package main

import (
	"advent-of-code-2025/day1"
	"advent-of-code-2025/day2"
	"advent-of-code-2025/day3"
	"advent-of-code-2025/day4"
	"advent-of-code-2025/day5"
	"fmt"
	"time"
)

func main() {
	runBenchmark(day1.TITLE+" Part 1", func() any { return day1.SolvePart1(day1.PATH) })
	runBenchmark(day1.TITLE+" Part 2", func() any { return day1.SolvePart2(day1.PATH) })
	runBenchmark(day2.TITLE+" Part 1", func() any { return day2.SolvePart1(day2.PATH) })
	runBenchmark(day2.TITLE+" Part 2", func() any { return day2.SolvePart2(day2.PATH) })
	runBenchmark(day3.TITLE+" Part 1", func() any { return day3.SolvePart1(day3.PATH) })
	runBenchmark(day3.TITLE+" Part 2", func() any { return day3.SolvePart2(day3.PATH) })
	runBenchmark(day4.TITLE+" Part 1", func() any { return day4.SolvePart1(day4.PATH) })
	runBenchmark(day4.TITLE+" Part 2", func() any { return day4.SolvePart2(day4.PATH) })
	runBenchmark(day5.TITLE+" Part 1", func() any { return day5.SolvePart1(day5.PATH) })
}

func runBenchmark(title string, solve func() any) {
	start := time.Now()
	solution := solve()
	elapsed := time.Since(start)
	fmt.Printf("%s: [duration:%s] [solution:%d]\n", title, elapsed, solution)
}
