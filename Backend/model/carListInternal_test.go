package model

import "testing"

func Test_prettyPrice(t *testing.T) {
	price := 1000000000
	want := "1 000 000 000 р"
	if got := prettyPrice(price); got != want {
		t.Errorf("prettyPrice() = %v, want %v", got, want)
	}
}
