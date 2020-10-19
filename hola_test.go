package main

import "testing"

func TestHola(t *testing.T) {
	// Step1 : test for empty argument
	emptyResult := hola("")

	if emptyResult != "Hola Dude!" {
		t.Errorf("hola(\"\") failed, expected %v, got %v", "Hola Dude!", emptyResult)
	}

	// Step2 : test for valid argument
	result := hola("Mike")

	if result != "Hola Mike!" {
		t.Errorf("hola(\"Mike\") failed, expected %v, got %v", "Hola Mike!", result)
	}
}
