// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
    var s string
    start := time.Now()
    for i := 0; i < len(os.Args); i++ {
        s += fmt.Sprintf("Arg%v: %v\n", i, os.Args[i])
    }
    fmt.Println(s)
    fmt.Printf("Elapsed time printing using for loop: %.4fs.\n", time.Since(start).Seconds())

    start = time.Now()
    s = strings.Join(os.Args, "\n") 
    fmt.Println(s)
    fmt.Printf("Elapsed time printing using join: %.4fs.\n", time.Since(start).Seconds())
}