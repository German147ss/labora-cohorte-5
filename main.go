package main

import "fmt"

func main() {

	//PONIENDO A PRUEBA EL SCAN
	/* var edad int
	fmt.Scan(&edad) */

	var edad2 int
	fmt.Scan(&edad2)

	//fmt.Scan(edad2): En este caso, estoy pasando la variable edad2 directamente, sin su dirección de memoria. Esto no funcionará como se espera. fmt.Scan requiere un puntero (dirección de memoria) para poder modificar la variable que le paso :(. Si no uso un puntero, fmt.Scan no podrá almacenar el valor leído en edad2.

	fmt.Println("La edad 2:", edad2)
}
