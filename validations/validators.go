package validations

import (
	"fmt"
	"strconv"
)

func AtoiCheck(str string) int {
	switch str {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	case "D":
		return 4
	case "E":
		return 5
	case "F":
		return 6
	case "G":
		return 7
	case "H":

		return 8
	}
	return 0
}

func ValidateMove() (bool, int, int, int, int) {

	// VALUES
	var from string
	var to string
	allowleters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	answers := []bool{}
	answerssec := []bool{}
	// Scan FROM
	fmt.Println("Введите клетку ОТКУДА хотите переместить шашку:")
	fmt.Scan(&from)
	if len(from) == 0 {
		fmt.Println("Ошибка ввода: Пустая строка")
		return false, 0, 0, 0, 0
	}

	firstLeterFromInByte := from[0]
	firstLeterFrom := string(firstLeterFromInByte)
	lastLeterFromInByte := from[1]
	lastLeterFrom := string(lastLeterFromInByte)
	// ATOI
	Vfirst, _ := strconv.Atoi(firstLeterFrom)
	if Vfirst != 0 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Cначала буква, потом цифра")
		return false, 0, 0, 0, 0
	}

	Vsecond, _ := strconv.Atoi(lastLeterFrom)
	if Vsecond == 0 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Cначала буква, потом цифра")
		return false, 0, 0, 0, 0
	}
	//ATOICHECK
	firstvalue := AtoiCheck(firstLeterFrom)
	if firstvalue == 0 {
		fmt.Println("Ошибка конвертации: Неправильно задана строка")
		return false, 0, 0, 0, 0
	}
	if len(from) >= 3 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Превышен лимит символов")
		return false, 0, 0, 0, 0
	}

	//FIRSTLETTER
	for _, i := range allowleters {
		if firstLeterFrom != i {
			answers = append(answers, false)
		} else {
			answers = append(answers, true)

		}
	}
	var falseCount int
	for _, v := range answers {
		if v {

		} else {
			falseCount++
		}
	}
	if falseCount == 8 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Неправильно введена буквенная колонка")
		return false, 0, 0, 0, 0
	}
	//SECOMDLETTER

	if Vsecond < 1 || Vsecond > 8 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Неправильно введена числовая колонка")
	}

	//TO
	fmt.Println("Введите клетку КУДА хотите переместить шашку:")
	fmt.Scan(&to)
	if len(to) == 0 {
		fmt.Println("Ошибка ввода: Пустая строка")
		return false, 0, 0, 0, 0
	}

	firstLeterToInByte := to[0]
	firstLeterTo := string(firstLeterToInByte)
	lastLeterToInByte := to[1]
	lastLeterTo := string(lastLeterToInByte)
	// ATOI
	Vthird, _ := strconv.Atoi(firstLeterTo)
	if Vthird != 0 {
		fmt.Println("Ошибка перемещения: Неправильный синтаксис. Cначала буква, потом цифра")
		return false, 0, 0, 0, 0
	}

	Vfourth, _ := strconv.Atoi(lastLeterTo)
	if Vsecond == 0 {
		fmt.Println("Ошибка перемещения: Неправильный синтаксис. Cначала буква, потом цифра")
		return false, 0, 0, 0, 0
	}
	//ATOICHECK
	thirdvalue := AtoiCheck(firstLeterTo)
	if thirdvalue == 0 {
		fmt.Println("Ошибка конвертации: Неправильно задана строка")
		return false, 0, 0, 0, 0
	}
	if len(to) >= 3 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Превышен лимит символов")
		return false, 0, 0, 0, 0
	}

	//FIRSTLETTER
	for _, i := range allowleters {
		if firstLeterTo != i {
			answerssec = append(answerssec, false)
		} else {
			answerssec = append(answerssec, true)

		}
	}
	var falseCountsec int
	for _, v := range answerssec {
		if v {

		} else {
			falseCountsec++
		}
	}
	if falseCountsec == 8 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Неправильно введена буквенная колонка")
		return false, 0, 0, 0, 0
	}
	//SECOMDLETTER

	if Vfourth < 1 || Vfourth > 8 {
		fmt.Println("Ошибка ввода: Неправильный синтаксис. Неправильно введена числовая колонка")
	}

	return true, Vsecond, firstvalue, Vfourth, thirdvalue
}
