package main

import (
	"fmt"
	"testing"
)

var list = []int{1721, 979, 366, 299, 675, 1456}

func TestGetFileOfInts(t *testing.T) {
	result, err := getFileOfInts("test/input")
	if err != nil {
		t.Fatal(err)
	}

	if !Equal(result, list) {
		fmt.Println(result)
		fmt.Println(list)
		t.Error("does not match")
	}
}

func TestGetMultiply2ThatAddTo2020(t *testing.T) {
	result, err := getFileOfInts("test/input")
	if err != nil {
		t.Fatal(err)
	}

	x, err := getMultiply2ThatAddTo2020(result)
	if err != nil {
		t.Fatal(err)
	}

	if x != 514579 {
		t.Fatalf("answer %d, expecting 514579", x)
	}
}

func TestGetMultiply3ThatAddTo2020(t *testing.T) {
	result, err := getFileOfInts("test/input")
	if err != nil {
		t.Fatal(err)
	}

	x, err := getMultiply3ThatAddTo2020(result)
	if err != nil {
		t.Fatal(err)
	}

	if x != 241861950 {
		t.Fatalf("answer %d, expecting 241861950", x)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
