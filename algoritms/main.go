package main

import (
	"fmt"
)

// FIBONACCI FUNCTION
func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// FACTORIAL FUNCTION
func factorial(n int64): int64 {
	
	if n == 1 {
		return 1
	}

	return n * factorial(n - 1)

}

// TCO: Tail Call Optimization
/*
- Os frames adicionais vão sendo removidos da pilha de execução.
- Recursividade seja tão eficiente quanto um loop.
- Algumas linguagens de programação não suportam TCO nativamente.

Técnica:
-[acc] é um acumulador (de intermediários) que vai sendo passado para a próxima chamada recursiva.
-[n] é o número que vai sendo decrementado a cada chamada recursiva.


*/

// MAIN FUNCTION TO RUN
func main() {
	//fmt.Println("Fatorial:", factorial(50000))
	fmt.Println("Fibonacci result:", fibonacci(10))
	fmt.Println("Factorial result": factorial(5))
}
