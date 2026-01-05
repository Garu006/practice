// CREAR LA SOLUCION PARA ESTO

package main

import (
	"fmt"
	"math"
	"time"
)

func SimularMovimiento(
	startRow, startCol int,
	instrucciones string,
	height, width int,
) ResultadoMovimiento{
	fila := startRow
	columna := startCol
	salioDelMapa := false
	pasosRealizados := 0

	for _, instruccion := range instrucciones {
		if salioDelMapa {
			break
		}
		switch instruccion {
		case 'U':
			fila -= 1
		case 'D':
			fila += 1
		case 'L':
			columna -= 1
		case 'R':
			columna += 1
		}
		pasosRealizados += 1

		if fila < 0 || fila >= height || columna < 0 || columna >= width {
			salioDelMapa = true
			// Ajustar fila y columna para que estén dentro del mapa
			fila = int(math.Max(0, math.Min(float64(fila), float64(height-1))))
			columna = int(math.Max(0, math.Min(float64(columna), float64(width-1))))
		}
	}
	return ResultadoMovimiento {
		filaFinal: fila,
		columnaFinal: columna,
		salioDelMapa: salioDelMapa,
		pasosRealizados: pasosRealizados,
	}
}

// ResultadoMovimiento debe contener: fila final, columna final, salió del mapa (bool), pasos realizados
// reglas 
// U: fila -1
// D: fila +1
// L: columna -1
// R: columna +1

type ResultadoMovimiento struct {
	filaFinal int
	columnaFinal int
	salioDelMapa bool
	pasosRealizados int
}

func main() {
	resultado := SimularMovimiento(2, 3, "UDLR", 5, 5)
	fmt.Println(resultado)

	resultado2 := SimularMovimiento(0, 0, "UUUU", 3, 3)
	fmt.Println(resultado2)

	resultado3 := SimularMovimiento(1, 1, "RRDDLLUU", 4, 4)
	fmt.Println(resultado3)

	resultado4 := SimularMovimiento(3, 3, "LLLLDDRRUU", 5, 5)
	fmt.Println(resultado4)

	resultado5 := SimularMovimiento(0, 0, "RRRRDDDDLLLLUUUU", 4, 4)
	fmt.Println(resultado5)

	resultado6 := SimularMovimiento(2, 2, "UUUURRRRDDDDLLLL", 5, 5)
	fmt.Println(resultado6)

	resultado7 := SimularMovimiento(1, 1, "UDUDUDUDUD", 3, 3)
	fmt.Println(resultado7)

	time.Sleep(1 * time.Second) // Pausa para ver los resultados antes de que termine el programa
}
