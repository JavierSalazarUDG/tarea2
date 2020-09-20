package main

import (
	"fmt"
	"math"
)

var (
	alumno    = "Reynaldo Javier Salazar Ochoa"
	tarea     = "Tarea 1"
	separador = "******************************************"
)

func main() {
	fmt.Println(tarea)
	fmt.Println(alumno)
	fmt.Println(separador)
	loop := false
	var ops int
	for loop == false {
		fmt.Println("Eliga una de las siguientes opciones:")
		fmt.Println("1-Calcular el área del cuadrado")
		fmt.Println("2-Calcular el área del triángulo")
		fmt.Println("3-Calcular el área del círculo")
		fmt.Println("4-Convertir de grados Fahrenheit a Celcius")
		fmt.Println("5-Salir")
		fmt.Print("-> ")
		fmt.Scan(&ops)
		switch ops {
		case 1:
			result := squareArea()
			fmt.Printf("El resultado es: %f\n", result)
		case 2:
			result := triangleArea()
			fmt.Printf("El resultado es: %f\n", result)
		case 3:
			result := circleArea()
			fmt.Printf("El resultado es: %f\n", result)
		case 4:
			result := covertToCelcius()
			fmt.Printf("El resultado es: %f\n", result)
		case 5:
			loop = true
		}
	}
}
func circleArea() float64 {

	r := 0.0
	fmt.Println("Calcula el Area de un circulo")
	fmt.Println("Ingrese el radio del circulo")
	fmt.Scan(&r)
	res := r * r
	return math.Pi * res
}
func squareArea() float64 {
	a := 0.0
	fmt.Println("Calcula el Area de un cuadrado")
	fmt.Println("Ingrese el valor de un lado del cuadrado")
	fmt.Scan(&a)
	return a * a
}
func triangleArea() float64 {
	b := 0.0
	a := 0.0

	fmt.Println("Calcula el Area de un trinagulo")
	fmt.Println("Ingrese el valor de la base")
	fmt.Scan(&b)
	fmt.Println("Muy bien Ahora ingrese la altura del tringulo")
	fmt.Scan(&a)
	return (b * a) / 2
}
func covertToCelcius() float64 {
	g := 0.0
	fmt.Println("Convertir grados Fahrenheit a Centígrados")
	fmt.Println("Ingrese grados Fahrenheit ")
	fmt.Scan(&g)
	return (g - 32) / 1.8000
}
