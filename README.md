Paquete Utilidades en Go

Este repositorio contiene un paquete en Go con dos funcionalidades principales:

Conversor de Monedas
Convierte un valor en USD a las siguientes monedas:

Euro (EUR)

Libra Esterlina (LB)

Won Surcoreano (WON)

Bitcoin (BTC)

Contador de Vocales
Dada una frase, cuenta cuántas veces aparece cada vocal (a, e, i, o, u).

Instalación del paquete

Clonar este repositorio:

git clone https://github.com/MasterPieceDR/paquete_utilidades.git


O, si se desea usarlo en otro proyecto, primero inicializar un módulo y luego importar el paquete:

go mod init proyecto_prueba
go get github.com/MasterPieceDR/paquete_utilidades

Uso en un proyecto

Ejemplo de main.go utilizando el paquete:

package main

import (
	"fmt"

	utilidades "github.com/MasterPieceDR/paquete_utilidades"
)

func main() {
	// Conversor de monedas
	var valor float64
	var opcion int
	fmt.Print("Ingresa el valor en USD: ")
	fmt.Scanln(&valor)
	fmt.Println("1) EUR, 2) LB, 3) WON, 4) BTC")
	fmt.Scanln(&opcion)
	fmt.Println("Convertido:", utilidades.Conversor(valor, opcion))

	// Contador de vocales
	var frase string
	fmt.Print("Ingresa una frase: ")
	fmt.Scanln(&frase)
	fmt.Println("Conteo de vocales:", utilidades.ContadorVocales(frase))
}

Estructura del proyecto
paquete_utilidades/
│── go.mod
│── utilidades.go
│── README.md

Ejemplo de salida

Ejecutando el programa de prueba (go run main.go):

Ingresa el valor en USD: 100
1) EUR, 2) LB, 3) WON, 4) BTC
1
Convertido: 92
Ingresa una frase: Hola Mundo desde Go
Conteo de vocales: a: 1, e: 2, i: 0, o: 3, u: 1

Autor

Diego Ruiz

Universidad Internacional del Ecuador (UIDE)
