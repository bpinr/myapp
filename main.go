package main

import (
	"errors"
	"fmt"
	"math/rand"
)

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

	for i := 0; i < 15; i++ {

		fmt.Println("Введите число")
		var num int //число пользовтеля
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		switch {
		case z < num:
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println("Верно")
			return
		}
	}

}

func play2() {
	z := rand.Intn(101)

	for i := 0; i < 10; i++ {
		fmt.Println("Введите число")
		var num int //число пользовтеля
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		switch {
		case z < num:
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println("Верно")
			return
		}
	}

}

func play3() {
	z := rand.Intn(201)

	for i := 0; i < 5; i++ {
		fmt.Println("Введите число")
		var num int //число пользовтеля
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		switch {
		case z < num:
			fmt.Println("Введённое значение слишком большое")
		case z > num:
			fmt.Println("Введённое значение слишком маленькое")
		case z == num:
			fmt.Println("Верно")
		}
	}

}

func playGame(a *int) {
	switch {
	case *a == 1:
		play1()
		fmt.Println("Начинается игра на уровне сложности 1")
	case *a == 2:
		play2()
		fmt.Println("Начинается игра на уровне сложности 2")
	case *a == 3:
		play3()
		fmt.Println("Начинается игра на уровне сложности 3")
	}
}

func main() {
	var a int //число сложности

	fmt.Println("Выберите сложность:")
	fmt.Println("Easy: 1 -50, 15 попыток")
	fmt.Println("Medium: 2 -100, 10 попыток")
	fmt.Println("Hard: 3 -200, 5 попыток")

	fmt.Scan(&a)

	_, err := difficult(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
}
