package main

import "testing"

func Test_normalize(t *testing.T) {
	tests := []struct {
		name   string
		msisdn string
		want   string
	}{
		{"start w zero", "08123456789", "628123456789"},
		{"start w plus and zero", "+08123456789", "628123456789"},
		{"start w plus and 62", "+628123456789", "628123456789"},
		{"start w 8", "8123456789", "628123456789"},
		{"start w 021", "02123456789", "622123456789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.msisdn); got != tt.want {
				t.Errorf("normalize(%v) = %v, want %v", tt.msisdn, got, tt.want)
			}
		})
	}
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalize("08123456789")
	}

}
