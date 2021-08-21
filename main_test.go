package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	input := "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM"
	firstRover, secondRover, err := Run(input)
	if err != nil {
		t.Fatal(err)
	}
	if firstRover.String() != "1 3 N" || secondRover.String() != "5 1 E" {
		t.Fatalf("expected value of first rover %s and expected value of second rover %s\n", "1 3 N", "5 1 E")
	}
}

func BenchmarkRun(b *testing.B) {
	input := "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM"
	for i := 0; i < b.N; i++ {
		Run(input)
	}
}
