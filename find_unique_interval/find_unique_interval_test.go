package main

import "testing"

func TestOverlapFunction(t *testing.T) {
	interval_a := interval{0, 4}
	interval_b := interval{3, 5}
	interval_c := interval{4, 5}
	interval_d := interval{6, 7}

	tests := []struct {
		name         string
		interval_lhs interval
		interval_rhs interval
		want         bool
	}{
		{
			name:         "lhs overlaps with rhs",
			interval_lhs: interval_a,
			interval_rhs: interval_b,
			want:         true,
		},
		{
			name:         "rhs is subset of lhs",
			interval_lhs: interval_b,
			interval_rhs: interval_c,
			want:         true,
		},
		{
			name:         "lhs is equal to rhs",
			interval_lhs: interval_a,
			interval_rhs: interval_a,
			want:         true,
		},
		{
			name:         "rhs does not overlap with lhs",
			interval_lhs: interval_a,
			interval_rhs: interval_d,
			want:         false,
		},
		{
			name:         "rhs does not overlap with lhs, but borders the interval",
			interval_lhs: interval_b,
			interval_rhs: interval_d,
			want:         false,
		},
	}

	for _, tt := range tests {
		has_overlap := overlaps(tt.interval_lhs, tt.interval_rhs)
		if has_overlap != tt.want {
			t.Errorf("Expected '%t', but got '%t'", tt.want, has_overlap)
		}
	}
}

func TestNaiveSearch(t *testing.T) {
	tests := []struct {
		name        string
		intervals   []interval
		want        bool
		want_result interval
	}{
		{
			name:        "first interval unique",
			intervals:   []interval{{4, 6}, {0, 3}, {7, 10}, {5, 7}},
			want:        true,
			want_result: interval{0, 3},
		},
		{
			name:        "last interval unique",
			intervals:   []interval{{25, 50}, {4, 6}, {7, 10}, {5, 7}},
			want:        true,
			want_result: interval{25, 50},
		},
		{
			name:        "interval that's not last or first is unique",
			intervals:   []interval{{3, 5}, {4, 6}, {7, 9}, {10, 30}, {10, 20}},
			want:        true,
			want_result: interval{7, 9},
		},
		{
			name:        "No unique overlap",
			intervals:   []interval{{1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:        false,
			want_result: interval{0, 0},
		},
	}

	for _, tt := range tests {
		result, res_interval := naive_search(tt.intervals)
		if result != tt.want {
			t.Errorf("Expected '%t', but got '%t'", tt.want, result)
		}
		if res_interval != tt.want_result {
			t.Errorf("Expected '[%d, %d]', but got '[%d, %d]'", tt.want_result.low, tt.want_result.high, res_interval.low, res_interval.high)

		}
	}
}

func TestDynamicSearch(t *testing.T) {
	tests := []struct {
		name        string
		intervals   []interval
		want        bool
		want_result interval
	}{
		{
			name:        "first interval unique",
			intervals:   []interval{{4, 6}, {0, 3}, {7, 10}, {5, 7}},
			want:        true,
			want_result: interval{0, 3},
		},
		{
			name:        "last interval unique",
			intervals:   []interval{{25, 50}, {4, 6}, {7, 10}, {5, 7}},
			want:        true,
			want_result: interval{25, 50},
		},
		{
			name:        "interval that's not last or first is unique",
			intervals:   []interval{{3, 5}, {4, 6}, {7, 9}, {10, 30}, {10, 20}},
			want:        true,
			want_result: interval{7, 9},
		},
		{
			name:        "No unique overlap",
			intervals:   []interval{{1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:        false,
			want_result: interval{0, 0},
		},
	}

	for _, tt := range tests {
		result, res_interval := dynamic_search(tt.intervals)
		if result != tt.want {
			t.Errorf("Expected '%t', but got '%t'", tt.want, result)
		}
		if res_interval != tt.want_result {
			t.Errorf("Expected '[%d, %d]', but got '[%d, %d]'", tt.want_result.low, tt.want_result.high, res_interval.low, res_interval.high)

		}
	}
}
