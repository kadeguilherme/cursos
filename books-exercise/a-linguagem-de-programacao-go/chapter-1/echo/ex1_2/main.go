// Ex1_2 exibe os argumentos de linha de comando com índice
package main

import (
	"fmt"
	"os"
)

func main() {
	// Exibe cada argumento com seu índice
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d %s\n", i, os.Args[i])
	}
}
