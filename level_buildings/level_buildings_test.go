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

import "testing"

func TestNaiveSearch(t *testing.T) {
	tests := []struct {
		name      string
		intervals []int
		want      int
	}{
		{
			name:      "one element",
			intervals: []int{1},
			want:      0,
		},
		{
			//            _
			//          _|x|
			//        _|x|x|
			//      _|o|o|o|
			//    _|x|o|o|o|
			//   |x|x|o|o|o|
			//    0 1 2 3 4
			//        S
			name:      "linear increase",
			intervals: []int{1, 2, 3, 4, 5},
			want:      6,
		},
		{
			//            _
			//           |x|
			//           |x|
			//           |x|
			//          _|x|
			//         |o|o|
			//         |o|o|
			//         |o|o|
			//         |o|o|
			//    _ _ _|o|o|
			//   |x|x|x|o|o|
			//    0 1 2 3 4
			//          S
			name:      "flat first",
			intervals: []int{1, 1, 1, 6, 10},
			want:      7,
		},
		{
			//     _ _ _ _ _
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//     0 1 2 3 4
			//     S
			name:      "equal height / pick first",
			intervals: []int{10, 10, 10, 10, 10},
			want:      0,
		},
		{
			//	           _
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//     _ _ _ _|o|
			//     0 1 2 3 4
			//             S
			name:      "Only one buidlin / pick lastg",
			intervals: []int{0, 0, 0, 0, 10},
			want:      0,
		},
	}

	for _, tt := range tests {
		result := level_buildings(tt.intervals, naive)
		if result != tt.want {
			t.Errorf("Expected '%d', but got '%d'", tt.want, result)
		}
	}
}

func TestDynamicSearch(t *testing.T) {
	tests := []struct {
		name      string
		intervals []int
		want      int
	}{
		{
			name:      "one element",
			intervals: []int{1},
			want:      0,
		},
		{
			//            _
			//          _|x|
			//        _|x|x|
			//      _|o|o|o|
			//    _|x|o|o|o|
			//   |x|x|o|o|o|
			//    0 1 2 3 4
			//        S
			name:      "linear increase",
			intervals: []int{1, 2, 3, 4, 5},
			want:      6,
		},
		{
			//            _
			//           |x|
			//           |x|
			//           |x|
			//          _|x|
			//         |o|o|
			//         |o|o|
			//         |o|o|
			//         |o|o|
			//    _ _ _|o|o|
			//   |x|x|x|o|o|
			//    0 1 2 3 4
			//          S
			name:      "flat first",
			intervals: []int{1, 1, 1, 6, 10},
			want:      7,
		},
		{
			//     _ _ _ _ _
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//    |o|o|o|o|o|
			//     0 1 2 3 4
			//     S
			name:      "equal height / pick first",
			intervals: []int{10, 10, 10, 10, 10},
			want:      0,
		},
		{
			//	           _
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//            |o|
			//     _ _ _ _|o|
			//     0 1 2 3 4
			//             S
			name:      "Only one buidlin / pick lastg",
			intervals: []int{0, 0, 0, 0, 10},
			want:      0,
		},
	}

	for _, tt := range tests {
		result := level_buildings(tt.intervals, dynamic)
		if result != tt.want {
			t.Errorf("Expected '%d', but got '%d'", tt.want, result)
		}
	}
}
