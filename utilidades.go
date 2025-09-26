package utilidades

import (
	"fmt"
	"strings"
)

func Conversor(valor float64, opcion int) float64 {
	var tasa float64

	switch opcion {
	case 1: // Euros
		tasa = 0.92
	case 2: // Libras Esterlinas
		tasa = 0.79
	case 3: // Won Surcoreano
		tasa = 1350.50
	case 4: // Bitcoin
		tasa = 0.000018
	default:
		tasa = 1
	}

	return valor * tasa
}

func ContadorVocales(frase string) string {
	frase = strings.ToLower(frase)

	// Contadores individuales para cada vocal
	a, e, i, o, u := 0, 0, 0, 0, 0

	// Recorrer cada letra de la frase
	for _, letra := range frase {
		switch letra {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		}
	}

	resultado := fmt.Sprintf("a: %d, e: %d, i: %d, o: %d, u: %d", a, e, i, o, u)

	return resultado
}
