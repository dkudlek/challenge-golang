/*
MIT License

# Copyright (c) 2022 David Kudlek

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type interval struct {
	low  int
	high int
}

func overlaps(lhs interval, rhs interval) bool {
	if rhs.low <= lhs.high && rhs.high >= lhs.low {
		return true
	}
	return false
}

// # Naive approach
//
// Complexity: O(N * N)
//
// Compare each intervale with all other intervals.
// Early exit when we find one interval that doesn't overlap with an other interval from the
// list.
func naive_search(list []interval) (bool, interval) {
	for idx, curr_interval := range list {
		has_overlap := false
		for second_idx, second_interval := range list {
			if idx == second_idx {
				continue
			} else if overlaps(curr_interval, second_interval) {
				has_overlap = true
				break // Early exit because we found an overlapping interval
			}
		}
		if !has_overlap {
			return true, curr_interval
		}

	}
	return false, interval{0, 0}
}

// # Dynamic approach
//  1. Sort the array: O(N log N)
//  2. Touch each element and compare to a memorized interval: O(N)
//
// Complexity: O(N*log(N)) + O(N) ~ O(N*log(N))
//
// Memoization technique: We use one interval to memorize all the intervals we've seen. When a
// interval overlaps with it then we grow this interval. This means for each element, we only
// need to compare against this interval. If it doesn't overlap then we create a new interval.
// If this is the last element or the next element does not overlap then we found an interval
// that doesn't overlap with any other interval.
// Early exit when we find one interval that doesn't overlap with an other interval from the
// list.
func dynamic_search(list []interval) (bool, interval) {
	// Sort list (and copy)
	sort.Slice(list, func(i, j int) bool {
		if list[i].low < list[j].low {
			return true
		} else if list[i].low == list[j].low && list[i].high <= list[j].high {
			return true
		}
		return false
	})
	sorted_list := list

	// Initialize other helper variables
	span := interval{0, 0}
	found := false
	idx_max := len(sorted_list) - 1
	for idx, itr := range sorted_list {
		// Update buffer and skip first check
		if idx == 0 {
			span = itr
			found = true
			continue
		}
		has_overlap := overlaps(span, itr)
		if has_overlap {
			if itr.high > span.high {
				span.high = itr.high
			}
			found = false
		} else {
			if idx == 1 {
				// First is single
				return true, span
			} else if idx == idx_max {
				// Last is single
				return true, itr
			} else if found {
				// Middle is single
				//
				// The last interval did't overlap with the temporary interval and the current
				// interval also doesn't overlap with the last one. The last one does not have
				// an overlap with any other interval.
				return true, span
			}
			span.low = itr.low
			span.high = itr.high
			found = true
		}
	}
	return false, interval{0, 0}
}

func to_time(duration time.Duration) string {
	hours := int(math.Mod((duration.Seconds() / 60.0), 60.0))
	minutes := int(math.Floor(duration.Seconds() / 60.0))
	seconds := int(math.Floor(duration.Seconds()))
	micros := int(duration.Microseconds() % 1000000)

	return fmt.Sprintf("%02d:%02d:%02d.%06d",
		hours,
		minutes,
		seconds,
		micros)
}

func readCsvFile(filePath string) []interval {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	list := []interval{}
	for idx, itr := range records {
		if idx == 0 {
			continue
		}
		low, err := strconv.Atoi(itr[0])
		if err != nil {
			log.Fatal("low value is not a number")
		}
		high, err := strconv.Atoi(itr[1])
		if err != nil {
			log.Fatal("high value is not a number")
		}
		list = append(list, interval{low, high})

	}

	return list
}

func execute_test(list []interval) {
	println("[RUN    ] Execute test: naive approach")
	naive_start := time.Now()
	naive_result, _ := naive_search(list)
	naive_duration := time.Since(naive_start)
	fmt.Printf(
		"[SUCCESS] Execute test: naive approach with '%v'\n",
		naive_result,
	)

	println("[RUN    ] Execute test: dynamic approach")

	dynamic_start := time.Now()
	dynamic_result, _ := dynamic_search(list)
	dynamic_duration := time.Since(dynamic_start)
	fmt.Printf(
		"[SUCCESS] Execute test: dynamic approach with '%v'\n",
		dynamic_result,
	)

	fmt.Printf(
		"[EVAL   ] Naive Approach took   %s || %12d us\n",
		to_time(naive_duration),
		naive_duration.Microseconds(),
	)
	fmt.Printf(
		"[EVAL   ] Dynamic Approach took %s || %12d us\n",
		to_time(dynamic_duration),
		dynamic_duration.Microseconds(),
	)

	println(naive_result == dynamic_result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func execute_random_test(n int) {
	// Random Test Suite
	println("[#######]")
	println("[RUN    ] Execute random test")
	max_size := int(math.Pow(2, 20))
	for i := 0; i < n; i++ {
		fmt.Printf("[RUN    ] Execute random test %d\n", i)

		list := []interval{}
		for j := 0; j < 1000000; j++ {
			one := rand.Intn(int(math.Pow(2, 32)))
			two := rand.Intn(int(math.Pow(2, 32)))
			if one > two {
				delta := abs(one-two) - max_size
				if delta > 0 {
					one = two + max_size
				}
				list = append(list, interval{two, one})
			} else {
				delta := abs(one-two) - max_size
				if delta > 0 {
					two = one + max_size
				}
				list = append(list, interval{one, two})
			}
		}
		execute_test(list)
	}
}

// # Find unique interval in list of intervals
//
// Given a list of intervals:
// We want to know if there's one interval which doesn't overlap with another interval
//
// An interval overlaps if end of one and start of the other are the equal (closed interval, including start and end value)
// e.g.
//   - "[0, 3]" and "[1, 2]" overlap
//   - "[0, 3]" and "[3, 5]" overlap
//   - "[0, 3]" and "[4, 6]" don't overlap
//
// # Solutions:
//  1. Naive Solution: "O(N * N)"
//  2. Dynamic solution: "O(N * log(N) + N) ~ O(N*log(N))"
//
// # Deliberations:
//   - tuple compare compares value by value:
//   - "(1, 2) < (2, 4)", because "1 < 2"
//   - "(1, 2) < (1, 3)", because "1 == 1 and 2 < 3"
//   - "(1, 2) > (0, 1)", because "1 > 0"
func main() {
	fmt.Println(os.Args)
	file_with_overlap := flag.String("file-with-overlap", "overlapping_intervals.csv", "")
	file_without_overlap := flag.String("file-without-overlap", "no_overlapping_intervals.csv", "")
	number_of_rand_runs := flag.Int("number-of-rand-runs", 0, "")
	flag.Parse()

	println("[#######]")
	println("[RUN    ] Test with overlap")
	overlap := readCsvFile(*file_with_overlap)
	execute_test(overlap)
	println("[SUCCESS] Test with overlap")
	println("[#######]")
	println("[RUN    ] Test without overlap")
	no_overlap := readCsvFile(*file_without_overlap)
	execute_test(no_overlap)
	println("[SUCCESS] Test without overlap")

	execute_random_test(*number_of_rand_runs)
}
