// Ex1_3 compara diferentes versões do programa echo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Medir versão ineficiente
	start := time.Now()
	result1 := ""
	for i := 1; i < len(os.Args); i++ {
		if i > 1 {
			result1 += " "
		}
		result1 += os.Args[i]
	}
	duration1 := time.Since(start)
	// Medir versão eficiente
	start = time.Now()
	strings.Join(os.Args[1:], " ")
	duration2 := time.Since(start)

	fmt.Printf("Versão ineficiente: %v\n", duration1)
	fmt.Printf("Versão eficiente: %v\n", duration2)
}
