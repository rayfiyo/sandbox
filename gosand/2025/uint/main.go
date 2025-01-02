package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ui uint = 18446744073709551615
	fmt.Printf("uint: %d\n      %d byte\n", ui, unsafe.Sizeof(ui))

	var i int = 9223372036854775807
	fmt.Printf(" int: %d\n      %d byte\n", i, unsafe.Sizeof(i))

	i = int(ui)

	fmt.Printf(" int: %d\n      %d byte\n", i, unsafe.Sizeof(i))
}
