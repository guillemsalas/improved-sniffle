package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findMaxPairSum(items []int) int {

	sort.Ints(items)

	sum, prev := 0, 0

	var positives []int

	for i, cur := range items {

		if cur < 0 {
			if prev == 0 {
				prev = cur
			} else {
				sum += prev * cur
				prev = 0
			}
		} else if cur == 0 {
			prev = 0
		} else if cur == 1 {
			sum += 1
		} else {
			positives = items[i:]
			break
		}
	}

	sum += prev
	prev = 0

	for i := len(positives) - 1; i >= 0; i-- {
		if prev != 0 {
			sum += prev * positives[i]
			prev = 0
		} else {
			prev = positives[i]
		}
	}
	return sum + prev
}

func parseFile(fileName string) []int {
	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	var items []int

	for scanner.Scan() {
		i, e := strconv.Atoi(scanner.Text())
		check(e)
		items = append(items, i)
	}
	return items
}

func processFile(fileName string) int {
	items := parseFile(fileName)
	return findMaxPairSum(items)
}

func main() {
	fileName := flag.String("file", "sample.txt", "source file")
	flag.Parse()

	result := processFile(*fileName)
	fmt.Println(result)
}