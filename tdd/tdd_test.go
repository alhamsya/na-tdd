package main

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{0}, false},
		{"13", args{13}, true},
		{"4", args{4}, false},
		{"104729", args{104729}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.number); got != tt.want {
				t.Errorf("isPrime(%d) = %v, want %v", tt.args.number, got, tt.want)
			}
		})
	}
}

func benchmarkPrime(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		primeGen(i)
	}
}

func BenchmarkPrime1(b *testing.B)    { benchmarkPrime(1, b) }
func BenchmarkPrime100(b *testing.B)  { benchmarkPrime(100, b) }
func BenchmarkPrime1000(b *testing.B) { benchmarkPrime(1000, b) }
