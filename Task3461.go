package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: go run Task3461.go input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Чтение содержимого входного файла
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении входного файла:", err)
		return
	}

	// Регулярное выражение для поиска простых математических выражений (например, 5+4=?, 5*4=?, 10/2=?)
	re := regexp.MustCompile(`(\d+)\s*([+\-*/])\s*(\d+)\s*=\?`)

	// Открываем файл для записи (очистим файл, если он существует)
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	// Чтение строк из входного файла
	lines := strings.Split(string(inputData), "\n")
	for _, line := range lines {
		// Поиск соответствий в строке
		match := re.FindStringSubmatch(line)
		if match != nil {
			// Преобразование строк в числа
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[3])
			operator := match[2]

			var result float64
			var err error

			// Обработка операций
			switch operator {
			case "+":
				result = float64(num1 + num2)
			case "-":
				result = float64(num1 - num2)
			case "*":
				result = float64(num1 * num2)
			case "/":
				if num2 == 0 {
					err = fmt.Errorf("деление на ноль")
				} else {
					result = float64(num1) / float64(num2)
				}
			}

			// Если ошибка деления на ноль, выводим сообщение
			if err != nil {
				_, _ = writer.WriteString(fmt.Sprintf("%s\n", err.Error()))
			} else {
				// Запись результата в выходной файл
				resultLine := fmt.Sprintf("%d%s%d=%.2f\n", num1, operator, num2, result)
				_, err := writer.WriteString(resultLine)
				if err != nil {
					fmt.Println("Ошибка при записи в выходной файл:", err)
					return
				}
			}
		}
	}

	// Сброс буфера в файл
	writer.Flush()
	fmt.Println("Результаты были записаны в", outputFile)
}
