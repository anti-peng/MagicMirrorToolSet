//usr/bin/env go run $0 "$@"; exit
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("== Hello Go Shell Script Example ==")
	for i, arg := range os.Args {
		fmt.Printf("arg[%d]='%s'\n", i, arg)
	}
}
