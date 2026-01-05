package strings

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},    // "b"
		{"pwwkew", 3},   // "wke"
		{"", 0},
		{" ", 1},
		{"au", 2},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
