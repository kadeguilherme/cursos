// Ex1_1 exibe os argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	// Exibe o nome do programa
	fmt.Println(os.Args[0])
}
