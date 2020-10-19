package main

import "testing"

func TestEmptyArgument(t *testing.T) {
	emptyResult := hola("")

	if emptyResult != "Hola Dude!" {
		t.Errorf("hola(\"\") failed, expected %v, got %v", "Hola Dude!", emptyResult)
	} else {
		t.Logf("hola(\"\") success, expected %v, got %v", "Hola Dude!", emptyResult)
	}
}

func TestValidArgument(t *testing.T) {
	result := hola("Mike")

	if result != "Hola Mike!" {
		t.Errorf("hola(\"Mike\") failed, expected %v, got %v", "Hola Mike!", result)
	} else {
		t.Logf("hola(\"Mike\") success, expected %v, got %v", "Hola Mike!", result)
	}
}
