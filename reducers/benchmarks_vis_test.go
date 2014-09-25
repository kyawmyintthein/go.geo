package reducers_test

import (
	"testing"

	reducers "."
)

func TestVisvalingamBenchmarkData(t *testing.T) {
	type reduceTest struct {
		Threshold float64
		Length    int
	}

	tests := []reduceTest{
		reduceTest{0.1, 875},
		reduceTest{0.5, 408},
		reduceTest{1.0, 291},
		reduceTest{1.5, 243},
		reduceTest{2.0, 210},
		reduceTest{3.0, 166},
		reduceTest{4.0, 149},
		reduceTest{5.0, 132},
	}
	path := benchmarkData()
	for i := range tests {
		p := reducers.VisvalingamThreshold(path, tests[i].Threshold)
		if p.Length() != tests[i].Length {
			t.Errorf("visvalingam benchmark data reduced poorly, got %d, expected %d", p.Length(), tests[i].Length)
		}
	}
}

func BenchmarkVisvalingamThreshold(b *testing.B) {
	path := benchmarkData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reducers.VisvalingamThreshold(path, 0.1)
	}
}

func BenchmarkVisvalingamKeep(b *testing.B) {
	path := benchmarkData()

	toKeep := int(float64(path.Length()) / 1.616)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reducers.VisvalingamKeep(path, toKeep)
	}
}
