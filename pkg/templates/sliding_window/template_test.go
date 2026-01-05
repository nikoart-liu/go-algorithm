package slidingwindow

import "testing"

func TestSlidingWindow(t *testing.T) {
	s := "abcabcbb"
	// 最长无重复子串是 "abc"，长度为 3
	expected := 3
	if got := SlidingWindow(s); got != expected {
		t.Errorf("SlidingWindow() = %v, want %v", got, expected)
	}

	s2 := "bbbbb"
	// 最长无重复子串是 "b"，长度为 1
	if got := SlidingWindow(s2); got != 1 {
		t.Errorf("SlidingWindow() = %v, want 1", got)
	}
}
