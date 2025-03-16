package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-reloaded/proces"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}

	// Чтение входного файла
	text, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error chteniya vkhodnogo fayla:", err)
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Printf("Error: vhodny faile dolzhny raznymi ")
	} else if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Error: Vkhodnoy fayl '%s' ne sushchestvuyet\n", inputFile)
	} else if filepath.Ext(inputFile) != ".txt" {
		fmt.Println("Error: Vkhodnoy fayl dolzhen imet' rasshireniye .txt ")
		os.Exit(1)
	} else if filepath.Ext(outputFile) != ".txt" {
		fmt.Println("Error: Vykhodnoy fayl dolzhen imet' rasshireniye .txt ")
		return
	}

	// Разделение текста на строки
	lines := strings.Split(string(text), "\n")
	var resultLines []string

	// Обработка каждой строки
	for _, line := range lines {
		processedLine := proces.ProcessLine(line)
		resultLines = append(resultLines, processedLine)
	}

	// Сборка результата в одну строку
	finalOutput := strings.Join(resultLines, "\n")

	// Запись результата в выходной файл
	err = os.WriteFile(os.Args[2], []byte(finalOutput), 0o644)
	if err != nil {
		fmt.Println("Error zapis' vykhodnogo fayla:", err)
	}
}
