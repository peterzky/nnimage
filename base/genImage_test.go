package base

import (
	"testing"
)

func TestZero(t *testing.T) {
	v := zero(3)
	if v != "000" {
		t.Error("Expected 000, got ", v)
	}
}

func Test(t *testing.T) {
	v := SeqRename(10)
	if v != "0010" {
		t.Error("Expected 0010, got ", v)
	}
}
