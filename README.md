# Symdiff: Finding the Symmetrical Difference Between Two Go Types

symdiff is a convenient library to help find the symmetrical difference between two similar types.

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/bacongobbler/symdiff"
)

type MyStruct struct {
	A string
	B int
}

func main() {
	a := MyStruct{
		A: "hello",
		B: 10,
	}
	b := MyStruct{
		A: "world",
		B: 10,
	}

	if err := symdiff.Diff(&a, b); err != nil {
		log.Fatal(err)
	}
	fmt.Println(a) // Output: {hello 0}
}
```
