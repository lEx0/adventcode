package main

import "testing"

func BenchmarkGetContainNumberFromStringLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetContainNumberFromString("seventhree1eightztszfourfivesix")
	}
}

func BenchmarkGetContainNumberFromStringShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetContainNumberFromString("8vrvqdrcmx9")
	}
}

func TestGetContainNumberFromString(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			"1abc2",
			12,
		},
		{
			"pqr3stu8vwx",
			38,
		},
		{
			"a1b2c3d4e5f",
			15,
		},
		{
			"treb7uchet",
			77,
		},
		{
			"two1nine",
			29,
		},
		{
			"eightwothree",
			83,
		},
		{
			"abcone2threexyz",
			13,
		},
		{
			"xtwone3four",
			24,
		},
		{
			"4nineeightseven2",
			42,
		},
		{
			"zoneight234",
			14,
		},
		{
			"7pqrstsixteen",
			76,
		},
		{
			"sevenine",
			79,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.input, func(t *testing.T) {
				if got := GetContainNumberFromString(tt.input); got != tt.want {
					t.Errorf("GetContainNumberFromString(%v) = %v, want %v", tt.input, got, tt.want)
				}
			},
		)
	}
}
