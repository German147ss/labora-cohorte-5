//Descomponer un número en una serie de sumas con restricciones:

package main

import "fmt"

func main() {
	var x int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&x)

	s1 := min(x, 50)
	x -= s1

	s2 := min(x, 50)
	x -= s2

	s3 := min(x, 600)
	x -= s3

	s4 := min(x, 800)
	x -= s4

	s5 := x

	fmt.Printf("S1: %d, S2: %d, S3: %d, S4: %d, S5: %d\n", s1, s2, s3, s4, s5)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
