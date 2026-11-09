package util

import "testing"

func TestShortTokenLength(t *testing.T) {
	token := NewToken()
	if got := len(token.Short()); got != 12 {
		t.Errorf("Short() length = %d, want 12", got)
	}
}
