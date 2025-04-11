package main

import (
	"fmt"
	"go-reloaded/proces"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Проверяем, указано ли правильное количество аргументов
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Использование: go run . <input_file.txt> [output_file.txt]")
		return
	}

	// Получаем путь к входному файлу из аргументов командной строки
	inputFile := os.Args[1]

	// Проверяем, имеет ли входной файл расширение .txt
	if filepath.Ext(inputFile) != ".txt" {
		fmt.Println("Ошибка: Входной файл должен иметь расширение .txt.")
		return
	}

	// Определяем путь к выходному файлу
	var outputFile string
	if len(os.Args) == 3 {
		outputFile = os.Args[2]
	} else {
		// Имя выходного файла по умолчанию, если не указано
		outputFile = "result.txt"
	}

	// Проверяем, имеет ли выходной файл расширение .txt
	if filepath.Ext(outputFile) != ".txt" {
		fmt.Println("Ошибка: Выходной файл должен иметь расширение .txt.")
		return
	}

	// Проверяем, что входной и выходной файлы не совпадают
	if filepath.Clean(inputFile) == filepath.Clean(outputFile) {
		fmt.Println("Ошибка: Входной и выходной файлы не должны быть одинаковыми.")
		return
	}

	// Читаем входной файл
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Разделяем входной текст на строки
	lines := strings.Split(string(data), "\n")

	// Обрабатываем каждую строку
	for i, textLine := range lines {
		textLine = proces.ProcessTrans(textLine)  // Преобразуем регистр текста
		textLine = proces.ReplaceNums(textLine)   // Преобразуем числа
		textLine = proces.ProscessPuncQ(textLine) // Обрабатываем пунктуацию и кавычки
		textLine = proces.Articl(textLine)
		lines[i] = textLine // Обновляем строку обработанным текстом
	}

	// Объединяем обработанные строки обратно в одну с переносами строк
	outputText := strings.Join(lines, "\n")

	// Записываем обработанный текст в выходной файл
	err = os.WriteFile(outputFile, []byte(outputText), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Обработка текста завершена. Проверьте выходной файл:", outputFile)
}
