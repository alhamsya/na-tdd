package main

import "testing"

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{"0", 0, false},
		{"13", 13, true},
		{"4", 4, false},
		{"104729", 104729, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.number); got != tt.want {
				t.Errorf("isPrime(%d) = %v, want %v", tt.number, got, tt.want)
			}
		})
	}
}

func benchmarkPrime(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		primeGen(i)
	}
}

func BenchmarkPrime1(b *testing.B)     { benchmarkPrime(1, b) }
func BenchmarkPrime100(b *testing.B)   { benchmarkPrime(100, b) }
func BenchmarkPrime10000(b *testing.B) { benchmarkPrime(10000, b) }
