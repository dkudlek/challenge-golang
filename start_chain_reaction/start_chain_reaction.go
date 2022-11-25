// # MIT License
//
// Copyright (c) 2022 David Kudlek.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"sort"
)

type FunElement struct {
	idx          int64 // array idx
	next_element int64 // idx + 1
	fun_val      int64
}

func execute(fun_val []int64, next_ptr []int64) int64 {
	if len(fun_val) != len(next_ptr) {
		panic("Vector of fun values is not equal to vector of next_ptrs!")
	}
	list := []FunElement{}
	for i, fun := range fun_val {
		list = append(list, FunElement{idx: int64(i), next_element: next_ptr[i], fun_val: fun})
	}
	sort.Slice(list, func(lhs, rhs int) bool {
		if list[lhs].next_element == list[rhs].next_element {
			if list[lhs].fun_val <= list[rhs].fun_val {
				return true
			} else {
				return false
			}
		} else if list[lhs].next_element < list[rhs].next_element {
			return true
		} else {
			return false
		}
	})

	new_order := []int64{}
	for _, val := range list {
		new_order = append(new_order, val.idx)
	}
	// Reverse new_order
	for i, j := 0, len(new_order)-1; i < j; i, j = i+1, j-1 {
		new_order[i], new_order[j] = new_order[j], new_order[i]
	}
	var fun_sum int64 = 0
	for vec_id := 0; vec_id < len(list); vec_id++ {
		curr_id := new_order[vec_id]
		if vec_id == len(new_order)-1 {
			fun_sum += fun_val[curr_id]
			break
		}
		next_id := new_order[vec_id+1]
		if next_ptr[curr_id] == 0 {
			fun_sum = fun_sum + fun_val[curr_id]
			continue
		} else if next_ptr[curr_id] == next_ptr[next_id] {
			// drop or swap
			if fun_val[curr_id] >= fun_val[next_id] {
				// drop
				fun_sum = fun_sum + fun_val[curr_id]
				continue
			} else {
				// swap
				fun_sum = fun_sum + fun_val[next_id]
				fun_val[next_id] = fun_val[curr_id]
				continue
			}
		} else {
			// merge
			merge_candidate := next_ptr[curr_id] - 1
			fun_val[merge_candidate] = max(fun_val[merge_candidate], fun_val[curr_id])
			continue

		}
	}
	return fun_sum

}

func max(lhs int64, rhs int64) int64 {
	if lhs >= rhs {
		return lhs
	}
	return rhs
}

func main() {}
