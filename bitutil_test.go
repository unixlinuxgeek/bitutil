package bitutil

import (
	"crypto/sha256"
	"testing"
)

func TestPopCount(t *testing.T) {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("y"))

	if Diff(a, b) == 0 {
		t.Errorf("Количество различающихся битов %d и %d не должно быть 0 получено %d", a, b, Diff(a, b))
	}
}
func TestCount(t *testing.T) {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("x"))
	if Diff(a, b) != 0 {
		t.Errorf("Количество различающихся битов %d и %d должно быть 0 но не %d", a, b, Diff(a, b))
	}
}
