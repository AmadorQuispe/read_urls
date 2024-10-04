package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abre el archivo de texto
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Mapa para eliminar duplicados
	uniqueURLs := make(map[string]struct{})
	count := 0
	// Leer el archivo línea a línea, sin cargarlo entero en memoria
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()

		// Comprobación simple para mejorar rendimiento, aquí podría usar una expresión regular
		if strings.HasSuffix(url, ".html") && strings.Contains(url, "shop") {
			if _, exists := uniqueURLs[url]; !exists {
				uniqueURLs[url] = struct{}{} // una estructura vacía para ahorrar memoria
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	// Mostrar resultados
	fmt.Println("Número total de URLs que cumplen los criterios:", count)
	fmt.Println("Listado de URLs:")
	for url := range uniqueURLs {
		fmt.Println(url)
	}
}
