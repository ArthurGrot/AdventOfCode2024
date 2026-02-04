package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantLeft  int
		wantRight int
	}{
		{
			name:      "simple numbers",
			line:      "3   4",
			wantLeft:  3,
			wantRight: 4,
		},
		{
			name:      "larger numbers",
			line:      "12345   67890",
			wantLeft:  12345,
			wantRight: 67890,
		},
		{
			name:      "single digit",
			line:      "1   2",
			wantLeft:  1,
			wantRight: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := parseLine(tt.line)
			if gotLeft != tt.wantLeft || gotRight != tt.wantRight {
				t.Errorf("parseLine() = (%d, %d); want (%d, %d)", gotLeft, gotRight, tt.wantLeft, tt.wantRight)
			}
		})
	}
}

func TestZipLists(t *testing.T) {
	tests := []struct {
		name     string
		leftIds  []int
		rightIds []int
		want     [][2]int
	}{
		{
			name:     "simple lists",
			leftIds:  []int{1, 2, 3},
			rightIds: []int{4, 5, 6},
			want:     [][2]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			name:     "single element",
			leftIds:  []int{10},
			rightIds: []int{20},
			want:     [][2]int{{10, 20}},
		},
		{
			name:     "empty lists",
			leftIds:  []int{},
			rightIds: []int{},
			want:     [][2]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zipLists(tt.leftIds, tt.rightIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zipLists() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestGetCopyOfList(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			name: "simple list",
			list: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "single element",
			list: []int{42},
			want: []int{42},
		},
		{
			name: "empty list",
			list: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCopyOfList(tt.list)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCopyOfList() = %v; want %v", got, tt.want)
			}
			// Verify it's a copy, not the same reference
			if len(tt.list) > 0 && &got[0] == &tt.list[0] {
				t.Error("getCopyOfList() returned same reference, not a copy")
			}
		})
	}
}

func TestSumOfAllDiffsPerPair(t *testing.T) {
	tests := []struct {
		name             string
		zippedListOfIds  [][2]int
		want             int
	}{
		{
			name:            "all positive diffs",
			zippedListOfIds: [][2]int{{1, 4}, {2, 5}, {3, 6}},
			want:            9, // |4-1| + |5-2| + |6-3| = 3 + 3 + 3 = 9
		},
		{
			name:            "mixed diffs",
			zippedListOfIds: [][2]int{{5, 2}, {10, 8}},
			want:            5, // |2-5| + |8-10| = 3 + 2 = 5
		},
		{
			name:            "zero diffs",
			zippedListOfIds: [][2]int{{5, 5}, {10, 10}},
			want:            0,
		},
		{
			name:            "single pair",
			zippedListOfIds: [][2]int{{1, 10}},
			want:            9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfAllDiffsPerPair(tt.zippedListOfIds)
			if got != tt.want {
				t.Errorf("sumOfAllDiffsPerPair() = %d; want %d", got, tt.want)
			}
		})
	}
}

func TestGetDictWithOccurrences(t *testing.T) {
	tests := []struct {
		name         string
		copyRightIds []int
		want         map[int]int
	}{
		{
			name:         "unique numbers",
			copyRightIds: []int{1, 2, 3, 4, 5},
			want:         map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{
			name:         "with duplicates",
			copyRightIds: []int{1, 2, 2, 3, 3, 3},
			want:         map[int]int{1: 1, 2: 2, 3: 3},
		},
		{
			name:         "all same number",
			copyRightIds: []int{5, 5, 5, 5},
			want:         map[int]int{5: 4},
		},
		{
			name:         "empty list",
			copyRightIds: []int{},
			want:         map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDictWithOccurrences(tt.copyRightIds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDictWithOccurrences() = %v; want %v", got, tt.want)
			}
		})
	}
}
