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
 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX",
 80: "LXXX", 90: "XC", 100: "C",
}

// Арабские числа
var arabicNumerals = map[string]int{
 "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
 "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
 "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
 "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
 "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80,
 "XC": 90, "C": 100,
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
 panic(fmt.Sprintf("Некорректное римское число: %s", roman))
}

// Перевод арабских чисел в римские
func arabicToRoman(num int) string {
 if num < 1 || num > 100 {
  panic(fmt.Sprintf("Результат должен быть в диапазоне от 1 до 100 для римских чисел"))
 }

 if num <= 20 {
  return romanNumerals[num]
 }

 tens := (num / 10) * 10
 units := num % 10

 if units == 0 {
  return romanNumerals[tens]
 }

 return romanNumerals[tens] + romanNumerals[units]
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
  panic(fmt.Sprintf("Некорректная операция: %s", op))
 }
}

// Основная функция
func main() {
 reader := bufio.NewReader(os.Stdin)
 fmt.Println("Введите выражение с помощью арабских или римских цифр (например, '3 + 5' или 'II + III'):")

 for {
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  // Разбиваем строку на части
  parts := strings.Fields(input)
  if len(parts) != 3 {
   panic("Некорректный ввод. Пожалуйста, введите выражение в формате 'число операция число'.")
  }

  aStr, op, bStr := parts[0], parts[1], parts[2]

  var a, b int
  var isRomanInput bool

  if isRoman(aStr) && isRoman(bStr)  {
   a = romanToArabic(aStr)
   b = romanToArabic(bStr)
   if a <= 10 && b <= 10 {
       isRomanInput = true
   } else {
    panic("Некорректный ввод чисел.")
  }
  } else if numA, err := strconv.Atoi(aStr); err == nil && numA >= 1 && numA <= 10 {
   a = numA
   if numB, err := strconv.Atoi(bStr); err == nil && numB >= 1 && numB <= 10 {
    b = numB
   } else {
    panic("Некорректное второе число.")
   }
  } else {
   panic("Некорректный ввод чисел.")
  }

  result := calculate(a, b, op)

  if isRomanInput {
   if result < 1 || result > 100 {
    panic(fmt.Sprintf("Результат вне диапазона для римских чисел: %d", result))
   }
   romanResult := arabicToRoman(result)
   fmt.Println("Результат:", romanResult)
  } else {
   fmt.Println("Результат:", result)
  }
 }
}
