package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)

	var height, width int
	var startCol, startRow int
	var instrucciones string

	fmt.Print("Ingrese el alto del mapa: ")
	fmt.Scanln(&height)
	fmt.Print("Ingrese el ancho del mapa: ")
	fmt.Scanln(&width)
	fmt.Print("Ingrese la fila inicial: ")
	fmt.Scanln(&startRow)
	fmt.Print("Ingrese la columna inicial: ")
	fmt.Scanln(&startCol)
	fmt.Print("Ingrese las instrucciones (U, D, L, R): ")
	instrucciones, _ = reader.ReadString('\n')
	instrucciones = strings.TrimSpace(instrucciones)

	// llamar a la funcion con los datos ingresados
	resultado, err := SimularMovimiento(startRow, startCol, instrucciones, height, width)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Resultado del movimiento: ")
	fmt.Println("Fila final: ", resultado.filaFinal)
	fmt.Println("Columna final: ", resultado.columnaFinal)
	fmt.Println("Salió del mapa: ", resultado.salioDelMapa)
	fmt.Println("Pasos realizados: ", resultado.pasosRealizados)
}
