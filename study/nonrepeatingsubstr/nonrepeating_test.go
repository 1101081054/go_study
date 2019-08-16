package main

import "testing"

func TestSubstr(t *testing.T){
	tests := []struct{
		s string
		ans int
	}{
		{"qwewertagj你好呀你好呀l", 10},
	}

	for _, tt := range tests {
		if length := lengthOfNonRepeatingSubStr(tt.s); length != tt.ans{
			t.Errorf("lengthOfNonRepeatingSubStr(%s) got %d; expected %d", tt.s, length, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s, ans := "qwewertagj你好呀你好呀l黑化肥会挥发", 13
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if length := lengthOfNonRepeatingSubStr(s); length != ans{
			b.Errorf("lengthOfNonRepeatingSubStr(%s) got %d; expected %d", s, length, ans)
		}
	}
}