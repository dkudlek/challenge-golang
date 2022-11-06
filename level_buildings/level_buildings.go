/**
 * MIT License
 *
 * Copyright (c) 2022 David Kudlek
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

/*
 * Given a list of building hights.
 * Select the building hight that minimizes the number of removed levels.
 * For each building smaller than the selected building, we remove all levels.
 * For each building higher than the selected building, we remove all levels above the selected building height.
 *
 * Input: [1, 2, 3, 4, 5]
 *          _
 *        _|x|                        S - Selected building height
 *      _|x|x|            _ _ _       x - Demolish levels
 *    _|o|o|o|     ->    |o|o|o|      o - keep levels
 *  _|x|o|o|o|           |o|o|o|
 * |x|x|o|o|o|        _ _|o|o|o|
 *  0 1 2 3 4
 *      S
 *
 * Return minimal number of demolished levels
 *
 * Trival:
 * - For each building height, test result: O(N * (N-1))
 *
 * Better solution:
 * - Sort the building heights: O(N log N)
 * - Memoization: Integrate building height over the list: O(N)
 * - For each possible building height, we can check the number of levels by taking the
 *     - left value: number of levels up to the selected building which we need to remove
 *     - last value: Total number of levels of all builds where substract the left value and the height of the selected buildings times the number of building after the selected building
 * - Walk over all levels and pick the best: O(N)
 *
 * Total: O(N log N) + O(N) + O(N) = O(N log N)
 *
 * Notes:
 * - input: [1, 2, 3, 4, 5]
 * - integrated: [1, 3, 6, 10, 15]
 * - Select idx: 2, height: 3
 * - val[idx-1] = 3 Levels need to be removed
 * - val[end] = 15 - val[idx] - (len(integrated) - 1 - idx ) * input[idx]
 *            = 15 - 6 - (5 - 1 - 2) * 3 = 9 - (2 * 3) = 3
 *
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

var naive = 0
var dynamic = 1

/**
 * Naive approach with: O(N * N)
 *
 * Compare each intervale with all other intervals.
 * Early exit when we find one interval that doesn't overlap with an other interval from the
 * list.
 */
func naive_approach(list_of_building_heights []int) int {
	min_demolished_levels := -1
	for _, selected_height := range list_of_building_heights {
		demolished_levels := 0
		for _, height := range list_of_building_heights {
			if height >= selected_height {
				demolished_levels += (height - selected_height)
			} else {
				demolished_levels += height
			}
		}
		if min_demolished_levels == -1 || demolished_levels < min_demolished_levels {
			min_demolished_levels = demolished_levels
		}
	}
	return min_demolished_levels
}

/**
 * Dynamic approach with: O(N*log(N)) + O(N) + O(N) ~ O(N*log(N))
 * (1) Sort: O(N*log(N))
 * (2) Integrate: O(N)
 * (3) Find min: O(N)
 *
 */
func dynamic_approach(list_of_building_heights []int) int {
	// Sort
	sort.Slice(list_of_building_heights, func(i, j int) bool {
		if list_of_building_heights[i] < list_of_building_heights[j] {
			return true
		}
		return false
	})
	// Memoization
	integrate_building_heights := []int{}
	accu := 0
	for _, el := range list_of_building_heights {
		accu = accu + el
		integrate_building_heights = append(integrate_building_heights, accu)
	}
	min_demolished_levels := -1
	for idx, height := range list_of_building_heights {
		left_levels := 0
		if idx > 0 {
			left_levels = integrate_building_heights[idx-1]
		}
		last_idx := len(integrate_building_heights) - 1
		right_levels := integrate_building_heights[last_idx] - integrate_building_heights[idx] - ((last_idx - idx) * height)
		demolished_levels := left_levels + right_levels
		if min_demolished_levels == -1 || demolished_levels < min_demolished_levels {
			min_demolished_levels = demolished_levels
		}
	}
	return min_demolished_levels
}

func has_negative(list_of_building_heights []int) bool {
	for _, val := range list_of_building_heights {
		if val < 0 {
			return true
		}
	}
	return false
}

func level_buildings(list_of_building_heights []int, methode int) int {
	if len(list_of_building_heights) == 0 {
		println("[WARN   ] List is empty!")
		return -1
	} else if has_negative(list_of_building_heights) {
		println("[ERROR  ] Negative building heights! Invalid!")
		return -1
	} else if len(list_of_building_heights) == 1 {
		println("[INFO   ] Only one element!")
		return 0
	} else if methode == naive {
		return naive_approach(list_of_building_heights)

	} else if methode == dynamic {
		return dynamic_approach(list_of_building_heights)
	}
	return -1
}

func sanity_check() {
	println("[RUN    ] Sanity check")
	/*
	            _
	          _|x|
	        _|x|x|
	      _|o|o|o|
	    _|x|o|o|o|
	   |x|x|o|o|o|
	    0 1 2 3 4
	        S
	*/
	building_heights := []int{1, 2, 3, 4, 5}
	min_levels := 6
	println(min_levels == level_buildings(building_heights, naive))
	println(min_levels == level_buildings(building_heights, dynamic))
	/*
	            _
	           |x|
	           |x|
	           |x|
	          _|x|
	         |o|o|
	         |o|o|
	         |o|o|
	         |o|o|
	    _ _ _|o|o|
	   |x|x|x|o|o|
	    0 1 2 3 4
	          S
	*/
	println("[SUCCESS] Sanity check: Default test")
	building_heights = []int{1, 1, 1, 6, 10}
	min_levels = 7
	println(min_levels == level_buildings(building_heights, naive))
	println(min_levels == level_buildings(building_heights, dynamic))
	/*
	    _ _ _ _ _
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	   |o|o|o|o|o|
	    0 1 2 3 4
	    S
	*/
	println("[SUCCESS] Sanity check: non linear building heights successfull")
	building_heights = []int{10, 10, 10, 10, 10}
	min_levels = 0
	println(min_levels == level_buildings(building_heights, naive))
	println(min_levels == level_buildings(building_heights, dynamic))
	/*
	           _
	          |o|
	          |o|
	          |o|
	          |o|
	          |o|
	          |o|
	          |o|
	          |o|
	          |o|
	   _ _ _ _|o|
	   0 1 2 3 4
	           S
	*/
	println("[SUCCESS] Sanity check: equal building heights successfull (pick first)")

	building_heights = []int{0, 0, 0, 0, 10}
	min_levels = 0
	println(min_levels == level_buildings(building_heights, naive))
	println(min_levels == level_buildings(building_heights, dynamic))
	println("[SUCCESS] Sanity check: pick last")

	// Negative numbers
	println(level_buildings([]int{1, -1}, naive) == -1)
	println(level_buildings([]int{1, -1}, dynamic) == -1)
	println("[SUCCESS] Sanity check: Negative levels")

	// Empty list
	println(level_buildings([]int{}, naive) == -1)
	println(level_buildings([]int{}, dynamic) == -1)
	println("[SUCCESS] Sanity check: Empty list")

	// Single Element
	println(level_buildings([]int{1}, naive) == 0)
	println(level_buildings([]int{1}, dynamic) == 0)
	println("[SUCCESS] Sanity check: Single Element")
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

func execute_test(list []int) {
	println("[RUN    ] Execute test: naive approach")
	naive_start := time.Now()
	naive_result := naive_approach(list)
	naive_duration := time.Since(naive_start)
	fmt.Printf(
		"[SUCCESS] Execute test: naive approach with '%v'\n",
		naive_result,
	)

	println("[RUN    ] Execute test: dynamic approach")

	dynamic_start := time.Now()
	dynamic_result := dynamic_approach(list)
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

func execute_random_test(n int) {
	/* Random Test Suite */
	println("[#######]")
	println("[RUN    ] Execute random test")
	for i := 0; i < n; i++ {
		fmt.Printf("[RUN    ] Execute random test %d\n", i)

		list := []int{}
		for j := 0; j < 10000; j++ {
			list = append(list, rand.Intn(int(math.Pow(2, 32))))

		}
		execute_test(list)
	}
}

func readCsvFile(filePath string) []int {
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
	list := []int{}
	for idx, itr := range records {
		if idx == 0 {
			continue
		}
		val, err := strconv.Atoi(itr[0])
		if err != nil {
			log.Fatal("low value is not a number")
		}
		list = append(list, val)

	}

	return list
}

func main() {
	fmt.Println(os.Args)
	file := flag.String("file", "level_buildings.csv", "")
	number_of_rand_runs := flag.Int("number-of-rand-runs", 0, "")
	flag.Parse()

	sanity_check()
	println("[#######]")
	println("[RUN    ] Test with overlap")
	buildings := readCsvFile(*file)
	execute_test(buildings)

	execute_random_test(*number_of_rand_runs)
}
