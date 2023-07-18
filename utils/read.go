package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(name string) string {
	// Abrir el archivo en modo lectura
	if len(name) == 0 {
		return ""
	}

	file, err := os.Open("arboles/" + name)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return ""
	}
	defer file.Close() // Cerrar el archivo al finalizar la función

	// Crear un lector para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Leer línea por línea y mostrar el contenido
	line := ""
	for scanner.Scan() {
		line = line + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
	return line
}
