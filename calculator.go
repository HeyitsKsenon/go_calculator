package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Римские числа
var romanNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

// Арабские числа
var arabicNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Является ли число римским
func isRoman(num string) bool {
	_, exists := arabicNumerals[num]
	return exists
}

// Перевод римских чисел в арабские
func romanToArabic(roman string) int {
	if val, exists := arabicNumerals[roman]; exists {
		return val
	}
	panic("Некорректное римское число")
}

// Перевод арабских чисел в римские
func arabicToRoman(num int) string {
	if num < 1 {
		panic("Результат не может быть меньше 1 для римских чисел")
	}
	for i, v := range romanNumerals {
		if i == num {
			return v
		}
	}
	panic("Римское число вне диапазона")
}

// Функция выполняет математические действия
func calculate(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Некорректная операция")
	}
}

// Основная функция, из которой вызываются остальные
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение с помощью арабских или римских цифр (например, '3 + 5' или 'II + III'):")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Разбиваем строку на части
		parts := strings.Fields(input)
		if len(parts) != 3 {
			panic("Некорректный ввод.")
		}

		aStr, op, bStr := parts[0], parts[1], parts[2]

		var a, b int
		var isRomanInput bool

		if isRoman(aStr) && isRoman(bStr) {
			a = romanToArabic(aStr)
			b = romanToArabic(bStr)
			isRomanInput = true
		} else if numA, err := strconv.Atoi(aStr); err == nil && numA >= 1 && numA <= 10 {
			a = numA
			if numB, err := strconv.Atoi(bStr); err == nil && numB >= 1 && numB <= 10 {
				b = numB
			} else {
				panic("Некорректное второе число")
			}
		} else {
			panic"Некорректный ввод чисел")
		}

		result := calculate(a, b, op)

		if isRomanInput {
			fmt.Println("Результат:", arabicToRoman(result))
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
