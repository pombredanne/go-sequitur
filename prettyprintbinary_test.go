package sequitur

import (
	"bytes"
	"fmt"
)

var testBinary = []byte{0xfe, 0xff, 0xfd, 0xfe, 0xff, 1, 2, 3, 4, 5, 'a', 'b', 1, 2, 3, 4, 5, 'a', 'b'}

func ExamplePrettyPrintBinary() {

	g := Parse(testBinary)

	var output bytes.Buffer
	if err := g.PrettyPrint(&output); err != nil {
		panic(err)
	}

	fmt.Println(output.String())

	// Output:
	// 0 -> 1 \xfd 1 2 2
	// 1 -> \xfe \xff
	// 2 -> \x01 \x02 \x03 \x04 \x05 a b
}
