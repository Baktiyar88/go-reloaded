package main

import (
	"fmt"
	"os"
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
		fmt.Println("Error reading input file:", err)
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
		fmt.Println("Error writing output file:", err)
	}
}
