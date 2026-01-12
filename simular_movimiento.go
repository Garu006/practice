package main

import (
	"bufio"
	"fmt"
	"time"
)

func SimularMovimiento(
	startRow, startCol int,
	instrucciones string,
	height, width int,
) (ResultadoMovimiento, error){
	fila := startRow
	columna := startCol
	salioDelMapa := false
	pasosRealizados := 0

	if height <= 0 || width <= 0 {
		return ResultadoMovimiento{}, fmt.Errorf("dimensiones del mapa invalidas")
	}
	if startRow < 0 || startRow >= height || startCol < 0 || startCol >= width {
		return ResultadoMovimiento{}, fmt.Errorf("posicion inicial fuera del mapa")
	}

	for _, instruccion := range instrucciones {
		filaAnterior := fila
		columnaAnterior := columna

		if instruccion != 'U' && instruccion != 'D' && instruccion != 'L' && instruccion != 'R' {
			return ResultadoMovimiento{}, fmt.Errorf("instruccion invalida: %c", instruccion)
		}

		switch instruccion {
		case 'U':
			fila--
		case 'D':
			fila++
		case 'L':
			columna--
		case 'R':
			columna++
		}
		pasosRealizados++

		if fila < 0 || fila >= height || columna < 0 || columna >= width {
			// Volver al ultimo estado valido
			fila = filaAnterior
			columna = columnaAnterior
			salioDelMapa = true
			break
		}
	}
	return ResultadoMovimiento {
		filaFinal: fila,
		columnaFinal: columna,
		salioDelMapa: salioDelMapa,
		pasosRealizados: pasosRealizados,
	}, nil
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
	resultado, err := SimularMovimiento(2, 3, "UDLR", 5, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return 
	}
	fmt.Println(resultado)

	resultado2, err := SimularMovimiento(0, 0, "UUUU", 3, 3)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado2)

	resultado3, err := SimularMovimiento(1, 1, "RRDDLLUU", 4, 4)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado3)

	resultado4, err := SimularMovimiento(3, 3, "LLLLDDRRUU", 5, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado4)

	resultado5, err := SimularMovimiento(0, 0, "RRRRDDDDLLLLUUUU", 4, 4)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado5)

	resultado6, err := SimularMovimiento(2, 2, "UUUURRRRDDDDLLLL", 5, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado6)

	resultado7, err := SimularMovimiento(1, 1, "UDUDUDUDUD", 3, 3)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resultado7)

	time.Sleep(1 * time.Second) // Pausa para ver los resultados antes de que termine el programa
}
