package main

import "testing"

type addTest struct {
	arg1, expected string
}

var addTests = []addTest{
	addTest{"Volvo released a new car with the following spec: V6 236HP. It will cost $22647 and going to be sold in New York only", "Volvo released a new car..."},
	addTest{"Volvo released", "Volvo released"},
	addTest{"Volvo released a new cars12 with the following spec: V6 236HP. It will cost $22647 and going to be sold in New York only", "Volvo released a new..."},
	addTest{"voLVO released a new ololloollo", "voLVO released a new..."},
	addTest{"voLVO released a new #############", "voLVO released a new..."},
	addTest{"$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$", "$$$$$$$$$$$$$$$$$$$$$$$$$..."},
	addTest{"", ""},
	addTest{" ", " "},
}

func TestShortTitle(t *testing.T) {

	for _, test := range addTests {
		if output := generateShortTitle(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
