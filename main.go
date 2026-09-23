package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Result struct {
	Date     string `json:"date"`
	Result   string `json:"result"`
	Attempts int    `json:"attempts"`
}

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func SaveResult(result string, attempts int) {
	file, err := os.OpenFile(
		"results.json",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer file.Close()

	data := Result{
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Result:   result,
		Attempts: attempts,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Ошибка создания JSON:", err)
		return
	}

	file.Write(jsonData)
	file.WriteString("\n")
}

func parseNumber(input string) (int, error) {
	return strconv.Atoi(input)
}

func hint(z int, num int) {

	r := z - num

	if r < 0 {
		r = -r
	}

	switch {
	case r <= 5:
		fmt.Println("🔥 Горячо")
	case r <= 15:
		fmt.Println("🙂 Тепло")
	default:
		fmt.Println("❄️ Холодно")
	}
}

func difficult(a *int) (int, error) {
	switch *a {
	case 1:
		play1()
		return 1, nil
	case 2:
		play2()
		return 2, nil
	case 3:
		play3()
		return 3, nil
	}
	return 0, errors.New("Введите верное число")
}

func play1() {
	z := rand.Intn(51)

	fmt.Println("Начинается игра на уровне сложности 1. Количество попыток: 15")
	var spisok []int
	for i := 0; i < 15; i++ {

		fmt.Println("Введите число")

		var input string
		fmt.Scan(&input)

		num, err := parseNumber(input)

		if err != nil {
			fmt.Println("Ошибка. Введите целое число!")
			i--
			continue
		}
		spisok = append(spisok, num)

		switch {
		case z < num:
			hint(z, num)
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			hint(z, num)
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println(Green + "Верно" + Reset)
			fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)

			SaveResult("win", i+1)

			return
		}

	}
	fmt.Println(Red + "Не угадал!Проигрыш" + Reset)
	fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)
	fmt.Println("Ответ был: ", z)

	SaveResult("lose", 15)

}

func play2() {
	z := rand.Intn(101)
	fmt.Println("Начинается игра на уровне сложности 2. Количество попыток: 10")
	var spisok []int
	for i := 0; i < 10; i++ {
		fmt.Println("Введите число")

		var input string
		fmt.Scan(&input)

		num, err := parseNumber(input)
		if err != nil {
			fmt.Println("Ошибка. Введите целое число!")
			i--
			continue
		}
		spisok = append(spisok, num)
		switch {
		case z < num:
			hint(z, num)
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			hint(z, num)
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println(Green + "Верно" + Reset)
			fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)

			SaveResult("win", i+1)

			return
		}
	}
	fmt.Println(Red + "Не угадал!Проигрыш" + Reset)
	fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)
	fmt.Println("Ответ был: ", z)

	SaveResult("lose", 10)

}

func play3() {
	z := rand.Intn(201)
	fmt.Println("Начинается игра на уровне сложности 3. Количество попыток: 5")
	var spisok []int
	for i := 0; i < 5; i++ {
		fmt.Println("Введите число")

		var input string
		fmt.Scan(&input)

		num, err := parseNumber(input)
		if err != nil {
			fmt.Println("Ошибка. Введите целое число!")
			i--
			continue
		}
		spisok = append(spisok, num)
		switch {
		case z < num:
			hint(z, num)
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			hint(z, num)
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println(Green + "Верно" + Reset)
			fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)

			SaveResult("win", i+1)

			return
		}
	}
	fmt.Println(Red + "Не угадал!Проигрыш" + Reset)
	fmt.Println(Yellow+"Вот твой список попыток:", spisok, Reset)
	fmt.Println("Ответ был: ", z)

	SaveResult("lose", 5)

}

func main() {

	for {
		var a int //число сложности

		fmt.Println("Выберите сложность:")
		fmt.Println("🟢Easy: 1 -50, 15 попыток")
		fmt.Println("🟡Medium: 2 -100, 10 попыток")
		fmt.Println("🔴Hard: 3 -200, 5 попыток")

		fmt.Scan(&a)

		_, err := difficult(&a)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Хотите продолжить игру? y/n")

		var s string
		fmt.Scan(&s)

		if s == "y" {
			continue
		}
		break
	}
}
