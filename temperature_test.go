package temperature_test

import (
	temperature "golang-temperature"
	"testing"
)

type TemperatureTest struct {
	name string
	input float64
	want float64
}

var cTemperatureTests = []TemperatureTest {
	{"test1", 50., 82.},
	{"test2", 0., 32.},
	{"test3", 20., 52.},
	{"test4", 500., 532.},
	{"test5", 5000., 5032.},
	{"test6", -50., -18.},
}

var fTemperatureTests = []TemperatureTest {
	{"test1", 82., 50.},
	{"test1", 83., 51.},
	{"test1", 182., 150.},
	{"test1", -190., -222.},
	{"test1", 0., -32.},
	{"test1", 32., 0.},
}

func runTestOnArray(name string, tests *[]TemperatureTest, t *testing.T, funcToTest func(tmp float64) float64) {
	for _,test := range *tests {
		got := funcToTest(test.input)
		if got != test.want {
			t.Errorf("TestCtoF(%s): input: %v; want: %v; got %v", test.name, test.input, test.want, got)
		}
	}
}

func TestCtoF(t *testing.T) {
	runTestOnArray("TestCtoF", &cTemperatureTests, t, temperature.CtoF)
}

func TestFtoC(t *testing.T) {
	runTestOnArray("TestFtoC", &fTemperatureTests, t, temperature.FtoC)
}

