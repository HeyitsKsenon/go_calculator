package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Римские числа
var romanNumerals = map[int]string {
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

// Арабские числа
var arabicNumerals = map[string]int {
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

	// Паника при некорректном римском число
	panic("Некорректное римское число")
}

// перевод арабских число в римские
func arabicToRoman(num int) string {
	// Паника при значении меньше одного
	if num < 1 {
		panic("Результат не может быть меньше 1 для римских чисел")
	}
	
	for i, v := range romanNumerals {
		if i == num {
			return v
	  }
	}

	// Паника при числе вне диапазона
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
	var input string
	// Выводим строку для взаимодействия с терминалом
	fmt.Println("Введите выражение с помощью арабских или римских цыфр (например, '3 + 5' или 'II + III'):")

	for {
		// Получаем введенные значения
		fmt.Scanln(&input)

		// Обрабатываем полученную строку 
		input = strings.TrimSpace(input)
		re := regexp.MustCompile(`^(\d+|[IVXLCDM]{1,3})(\s*[\+\-\*\/]\s*)(\d+|[IVXLCDM]{1,3})$`)
		matches := re.FindStringSubmatch(input)

		// Случай некорректного ввода данных
		if matches == nil {
			panic("Некорректный ввод. Попробуйте еще раз.")
		}

		// Распределяем числа и знак операции
		aStr, op, bStr := matches[1], matches[2], matches[3]

		var a, b int
		var isRomanInput bool

		// Проверяем какие числа пришли и корректны ли данные
		if isRoman(aStr) && isRoman(bStr) {
			a = romanToArabic(aStr)
			b = romanToArabic(bStr)
			isRomanInput = true
		} else if numA, ok := strconv.Atoi(aStr); ok == nil && numA >= 1 && numA <= 10 {
			a = numA

			if numB, ok := strconv.Atoi(bStr); ok == nil && numB >= 1 && numB <= 10 {
		   		b = numB
		  	} else {
		   		panic("Некорректное второе число")
		  	}
	 	} else {
		  	panic("Некорректный ввод чисел")
	 	}

		// Выполняем действие
	 	result := calculate(a, b, op)

		// Выводим результат в соответствующем формате
	 	if isRomanInput {
		  	fmt.Println("Результат:", arabicToRoman(result))
	 	} else {
		  	fmt.Println("Результат:", result)
	 	}
	}
}
