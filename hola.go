package main

import "fmt"

func hola(name string) string {
	if len(name) == 0 {
		return "Hola Dude!"
	} else {
		return fmt.Sprintf("Hola %v!", name)
	}
}
