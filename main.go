package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите математическое выражение: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		panic("Ошибка при чтении ввода: " + err.Error())
	}

	parts := strings.Fields(input)

	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	arg1, op, arg2 := parts[0], parts[1], parts[2]

	num1, num2, numType := getNumbers(arg1, arg2)

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("одно или несколько чисел выходят из допустимого диапазона")
	}

	calculation := calculate(num1, num2, op)

	if numType == "roman" {
		fmt.Println(convertArabToRoman(calculation))
	}
	if numType == "arabic" {
		fmt.Println(calculation)
	}

}

func getNumbers(arg1 string, arg2 string) (int, int, string) {

	romanToArabic := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	var num1, num2 int

	num1, isRoman1 := romanToArabic[arg1]
	num2, isRoman2 := romanToArabic[arg2]

	if isRoman1 != isRoman2 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if isRoman1 && isRoman2 {
		return num1, num2, "roman"
	}

	num1, _ = strconv.Atoi(arg1)
	num2, _ = strconv.Atoi(arg2)

	return num1, num2, "arabic"
}

func calculate(num1 int, num2 int, operator string) int {
	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	return result
}

func convertArabToRoman(arabNum int) string {
	if arabNum < 1 {
		panic("Выдача паники, так как в римской системе нет нуля и отрицательных чисел.")
	}

	arabicToRoman := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	var convResult string
	arabicValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, arabValue := range arabicValues {
		for arabNum >= arabValue {
			arabNum -= arabValue
			convResult += arabicToRoman[arabValue]
		}
	}

	return convResult
}
