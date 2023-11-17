// Suma y promedio de tres temperaturas:
package main

import "fmt"

func main() {
	var temp1, temp2, temp3 float64
	fmt.Print("Ingrese tres temperaturas: ")
	fmt.Scan(&temp1, &temp2, &temp3)

	suma := temp1 + temp2 + temp3
	promedio := suma / 3

	fmt.Printf("La suma es: %.2f y el promedio es: %.2f\n", suma, promedio)
}
