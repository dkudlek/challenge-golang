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
	"math"
	"testing"
)

func TestChainReaction(t *testing.T) {

	tests := []struct {
		name           string
		fun_values     []int64
		pointer_values []int64
		want           int64
	}{
		{
			name:           "single value",
			fun_values:     []int64{50},
			pointer_values: []int64{0},
			want:           50,
		},
		{
			name:           "Tow in a row",
			fun_values:     []int64{50, 40},
			pointer_values: []int64{0, 1},
			want:           50,
		},
		{
			name:           "Ten parallel",
			fun_values:     []int64{100, 90, 80, 70, 60, 50, 40, 30, 20, 10},
			pointer_values: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:           100,
		},
		{
			name:           "Multi subtree",
			fun_values:     []int64{100, 90, 80, 70, 60, 50, 40, 30, 20, 10},
			pointer_values: []int64{0, 1, 2, 3, 4, 0, 6, 7, 8, 9},
			want:           150,
		},
		{
			name:           "Mote Multi subtree",
			fun_values:     []int64{100, 90, 80, 70, 60, 50, 40, 30, 20, 10},
			pointer_values: []int64{0, 1, 0, 3, 0, 5, 0, 7, 0, 9},
			want:           300,
		},
		{
			name:           "Mote Multi subtree",
			fun_values:     []int64{100, 90, 80, 70, 60, 50, 40, 30, 20, 10},
			pointer_values: []int64{0, 1, 0, 3, 0, 5, 0, 7, 0, 9},
			want:           300,
		},
		{
			name:           "Execute merge",
			fun_values:     []int64{30, 40, 50, 60},
			pointer_values: []int64{0, 1, 1, 2},
			want:           110,
		},
		{
			name:           "Execute drop",
			fun_values:     []int64{30, 40, 50},
			pointer_values: []int64{0, 1, 1},
			want:           90,
		},
	}

	for _, tt := range tests {
		max_fun := execute(tt.fun_values, tt.pointer_values)
		if max_fun != tt.want {
			t.Errorf("Expected '%d', but got '%d'", tt.want, max_fun)
		}
	}
}

func TestMaxValue(t *testing.T) {
	max_val := int64(math.Pow(10, 9))
	fun_values := []int64{}
	pointer_values := []int64{}
	for i := 0; i < 1000; i++ {
		fun_values = append(fun_values, max_val)
		pointer_values = append(pointer_values, 0)
	}
	result := max_val * 1000
	max_fun := execute(fun_values, pointer_values)
	if max_fun != result {
		t.Errorf("Expected '%d', but got '%d'", result, max_fun)

	}
}

func TestPanic(t *testing.T) {
	// No need to check whether `recover()` is nil. Just turn off the panic.
	defer func() { _ = recover() }()

	execute([]int64{30, 40, 50, 60}, []int64{0, 1, 1})

	// Never reaches here if `OtherFunctionThatPanics` panics.
	t.Errorf("did not panic")
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		lhs  int64
		rhs  int64
		want int64
	}{
		{
			name: "lhs is bigger",
			lhs:  50,
			rhs:  40,
			want: 50,
		},
		{
			name: "rhs is bigger",
			lhs:  40,
			rhs:  50,
			want: 50,
		},
		{
			name: "equal",
			lhs:  50,
			rhs:  50,
			want: 50,
		},
	}

	for _, tt := range tests {
		max_val := max(tt.lhs, tt.rhs)
		if max_val != tt.want {
			t.Errorf("Expected '%d', but got '%d'", tt.want, max_val)
		}
	}
}
