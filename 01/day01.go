package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Abre o arquivo para leitura
	whichFile := func() string {
		if len(os.Args) < 2 {
			return "input.txt"
		}
		if os.Args[1] == "sample" {
			return "sample.txt"
		}
		return "input.txt"
	}()
	arquivo, err := os.Open(whichFile)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	sum := 0

	for scanner.Scan() {
		linha := scanner.Text()
		number := extractNumbers(linha)
		sum += number
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo:", err)
	}
}

func extractNumbers(input string) int {
	var firstDigit, lastDigit int

	for _, char := range input {
		if unicode.IsDigit(char) {
			if firstDigit == 0 {
				firstDigit = int(char) - 48
			}
			lastDigit = int(char) - 48
		}
	}

	return firstDigit * 10 + lastDigit
}