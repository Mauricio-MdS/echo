// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
    var s string
    for i := 0; i < len(os.Args); i++ {
        s += fmt.Sprintf("Arg%v: %v\n", i, os.Args[i])
    }
    fmt.Println(s)
}