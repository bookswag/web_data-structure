package utils

import "fmt"

func Push(item int) {

}

func Pop(item int) {

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func PrintPointer() {
	i := 1
	fmt.Println("init : ", i)

	zeroval(i)
	fmt.Println("input zero : ", i)

	zeroptr(&i)
	fmt.Println("input pointer to zero : ", i)

	fmt.Println("pointer value : ", &i)
}
