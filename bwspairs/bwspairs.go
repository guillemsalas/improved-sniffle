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

func pairArraySum(items []int) int {
	sort.Ints(items)

	sum, prev := 0, 0

	pivot := sort.SearchInts(items, 2)

	for _, cur := range items[:pivot] {
		switch {
		case cur < 0:
			if prev != 0 {
				sum += prev * cur
				prev = 0
			} else {
				prev = cur
			}
		case cur == 0:
			prev = 0
		case cur == 1:
			sum += 1
		}
	}
	sum += prev
	prev = 0

	items = items[pivot:]

	for i := len(items) - 1; i >= 0; i-- {
		if prev != 0 {
			sum += prev * items[i]
			prev = 0
		} else {
			prev = items[i]
		}
	}
	sum += prev

	return sum
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
	return pairArraySum(items)
}

func main() {
	fileName := flag.String("file", "sample.txt", "source file")
	flag.Parse()

	result := processFile(*fileName)
	fmt.Println(result)
}
