package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 7, 19},
	{"example2.txt", 5, 23},
	{"example3.txt", 6, 23},
	{"example4.txt", 10, 29},
	{"example5.txt", 11, 26}}

func TestTaskOne(t *testing.T) {

	for _, test := range testset {
		r := findsegment(test.fname, 4)
		if r != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		r := findsegment(test.fname, 14)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
