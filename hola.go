package main

import "fmt"

func hola(name string) string {
	if len(name) == 0 {
		return "Hola Dude!"
	}

	return fmt.Sprintf("Hola %v!", name)
}
