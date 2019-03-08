package temperature_test

import (
	temperature "golang-temperature"
	"testing"
)

func TestCtoF(t *testing.T) {
	want := 82.
	got := temperature.CtoF(50)
	if got != want {
		t.Fail()
	}
}

func TestFtoC(t *testing.T) {
	want := 50.
	got := temperature.FtoC(82)
	if got != want {
		t.Fail()
	}
}
