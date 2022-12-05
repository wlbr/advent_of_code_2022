package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 string
	expectedtask2 string
}

var testset []*testdata = []*testdata{{"example.txt", "CMZ", "MCD"}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		r := task1(test.fname)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%s' -  Wanted: '%s'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		r := task2(test.fname)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%s' -  Wanted: '%s'", test.fname, r, test.expectedtask2)
		}
	}
}
