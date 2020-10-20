package hola

import "fmt"

func Hola(name string) string {
	if len(name) == 0 {
		return "Hola Dude!"
	}

	return fmt.Sprintf("Hola %v!", name)
}
