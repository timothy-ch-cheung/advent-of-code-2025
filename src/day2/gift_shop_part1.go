package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type productRange struct {
	start string
	end   string
}

func SolvePart1(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	productRanges := parseLine(scanner.Text())
	var sum int64 = 0
	for i := 0; i < len(productRanges); i++ {
		sum += sumInvalidIds(productRanges[i].start, productRanges[i].end)
	}
	return sum
}

func parseLine(line string) []productRange {
	result := make([]productRange, 0)
	productRanges := strings.Split(line, ",")
	for i := 0; i < len(productRanges); i++ {
		parsedLine := strings.Split(productRanges[i], "-")
		start := parsedLine[0]
		end := parsedLine[1]
		result = append(result, productRange{start: start, end: end})
	}
	return result
}

func sumInvalidIds(start string, end string) int64 {
	lowerBound := findLowerBound(start)
	upperBound := findUpperBound(end)
	var total int64 = 0
	for i := lowerBound; i <= upperBound; i++ {
		half := toStr(i)
		total += toInt(half + half)
	}
	return total
}

func findLowerBound(start string) int64 {
	lowerBoundLength := findLowerBoundMinLength(start)
	if len(start)%2 == 1 {
		lowerBoundHalf := strconv.Itoa(pow(10, intDivide(len(start), 2)))
		return toInt(lowerBoundHalf)
	}

	startInt := toInt(start)
	lowerBoundHalf := start[:lowerBoundLength]
	lowerBound := toInt(lowerBoundHalf + lowerBoundHalf)
	if lowerBound < startInt {
		return toInt(lowerBoundHalf) + 1
	}
	return toInt(lowerBoundHalf)
}

func findLowerBoundMinLength(start string) int {
	startLength := len(start)
	half := intDivide(startLength, 2)
	if startLength%2 == 1 {
		return half + 1
	} else {
		return half
	}
}

func findUpperBoundMaxLength(end string) int {
	return intDivide(len(end), 2)
}

func findUpperBound(end string) int64 {
	upperBoundLength := findUpperBoundMaxLength(end)
	if len(end)%2 == 1 {
		return toInt(repeat("9", upperBoundLength))
	}
	endInt := toInt(end)
	upperBoundHalf := end[:upperBoundLength]
	upperBound := toInt(upperBoundHalf + upperBoundHalf)
	if upperBound <= endInt {
		return toInt(upperBoundHalf)
	}
	upperBoundHalf = toStr(toInt(upperBoundHalf) - 1)
	return toInt(upperBoundHalf)
}

func intDivide(numerator int, denominator int) int {
	remainder := numerator % denominator
	return (numerator - remainder) / denominator
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func repeat(str string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += str
	}
	return result
}

func toInt(str string) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("Error in toInt")
	}
	return result
}

func toStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
