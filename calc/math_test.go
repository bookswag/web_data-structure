package calc

import (
	"dsviewer/logger"
	"testing"
)

func TestMathAdd(t *testing.T) {
	_, result := Add(1, 2, 3)
	logger.PrintLogIntResult("Add(1, 2, 3)", result, 6, t)
}

func TestMathAddJustOne(t *testing.T) {
	_, result := Add(1)
	logger.PrintLogIntResult("Add(1", result, 0, t)
}
