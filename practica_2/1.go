/*
1. Realice un algoritmo para determinar si un número es perfecto. Un número es perfecto cuando la suma de sus divisores propios es igual al número. Los divisores propios de un número son todos sus divisores sin contar el mismo número.

Sean X1, X2, XN todos divisores propios de X

Si X es propio entonces X1 + x2 +…+ XN es igual a X

Ejemplo:

6 es un número perfecto

Divisores Propios: 1, 2 y 3.

1 + 2 + 3 = 6
*/

package main

import "fmt"

func esNumeroPerfecto(numero int) bool {
	sumaDivisores := 0

	for i := 1; i <= numero/2; i++ {
		if numero%i == 0 {
			sumaDivisores += i
		}
	}

	return sumaDivisores == numero
}

func main() {
	numero := 6
	if esNumeroPerfecto(numero) {
		fmt.Printf("%d es un número perfecto\n", numero)
	} else {
		fmt.Printf("%d no es un número perfecto\n", numero)
	}
}
