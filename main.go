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

	for {
		fmt.Println("Введите значение")
		text, err := reader.ReadString('\n') // Ждет ввода данных в формате строки
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
		}
		text = strings.TrimSpace(text)   // Удаляет начальные и конечные пробелы
		text = strings.ToUpper(text)     // Преобразует строку в верхний регистр
		strSlice := strings.Fields(text) // Разбивает строку на пробельные символы и возвращает срез, содержащий непробельные разделы строки

		if len(strSlice) == 1 { // Проверяет введенные данные на соответствие заданию
			fmt.Println("Ошибка. Cтрока не является математической операцией.")

		} else if len(strSlice) == 2 {
			fmt.Println("Ошибка. Cтрока не является математической операцией.")

		} else if len(strSlice) > 3 {
			fmt.Println("Ошибка. Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *): ")

		} else if strings.ContainsAny(strSlice[0], "12345678910") && strings.ContainsAny(strSlice[2], "IVX") || strings.ContainsAny(strSlice[2], "12345678910") && strings.ContainsAny(strSlice[0], "IVX") {
			fmt.Println("Ошибка. Используются одновременно разные системы исчисления")

		} else if strings.ContainsAny(strSlice[0], "12345678910") && strings.ContainsAny(strSlice[2], "12345678910") {
			a, err := strconv.Atoi(strSlice[0])
			if err != nil {
				fmt.Println("Ошибка преобразования первого числа: ", err)
			}

			b, err := strconv.Atoi(strSlice[2])
			if err != nil {
				fmt.Println("Ошибка преобразования второго числа: ", err)
			}

			if a > 10 || b > 10 || a < 1 || b < 1 {
				fmt.Println("Ошибка. Числа не удавлетворяют диапазону от 1 до 10")
			} else {

				fmt.Println(Arabic(a, b, strSlice[1]))
			}

		} else if strings.ContainsAny(strSlice[0], "IVX") && strings.ContainsAny(strSlice[2], "IVX") {

			if Decode(strSlice[0]) < Decode(strSlice[2]) || strings.ContainsAny(strSlice[0], "-") || strings.ContainsAny(strSlice[2], "-") {
				fmt.Println("Ошибка. В римской системе исчисления нет отрицательных чисел")

			} else {
				fmt.Println(Encode(Arabic(Decode(strSlice[0]), Decode(strSlice[2]), strSlice[1])))
			}

		}
	}
}

func Arabic(a, b int, s string) (res int) {

	switch s {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	return res
}

func Decode(roman string) (sum int) { // Конвертируем римские цифры в арабские
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

func Encode(num int) (roman string) { // Конвертируем арабские цифры в римские
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1

	}

	return roman
}
