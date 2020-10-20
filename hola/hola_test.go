package hola

import (
	"dsviewer/logger"
	"testing"
)

func TestEmptyArgument(t *testing.T) {
	emptyResult := Hola("")
	logger.PrintLogStringResult("hola(\"\")", emptyResult, "Hola Dude!", t)
}

func TestValidArgument(t *testing.T) {
	result := Hola("Mike")
	logger.PrintLogStringResult("hola(\"Mike\")", result, "Hola Mike!", t)
}
