package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var StrigConv strings.Builder // инициализируем строку возврата
	// исключительные случаи
	if len(s) == 0 {
		return StrigConv.String(), nil
	} // нулевая длина
	if s[0] >= '0' && s[0] <= '9' {
		return StrigConv.String(), ErrInvalidString
	} // первый символ - цифра

	// Для доп.задания - ловим \, предшествующий цифре (пока погодим)
	//  --

	// структура для разбора строки
	type razbor struct {
		ch        rune // символ
		type_simv int  // тип
		len       int  // кол-во повторений
	}
	rz := make([]razbor, len(s)) // количество элементов равно длине строки в байтах (а не в рунах)
	var i, type_prev_simv int    // счетчик, тип предыдущего символа

	for _, simv := range s {

		switch { // определяем тип текущего символа в строке
		case unicode.IsPrint(simv) && !unicode.IsDigit(simv):
			rz[i].type_simv = 'c' // печатный символ и не цифра
		case unicode.IsDigit(simv):
			rz[i].type_simv = 'd' // цифра
		//case unicode.IsLetter(t):  rz[i].ty='c'  // буква
		case unicode.IsControl(simv):
			rz[i].type_simv = 's' // спец.символ - к сожалению, слишком много
		default:
			{
				//rz[i].type_simv = '.'
				//pr_simv_no_dop = 1
				return StrigConv.String(), ErrInvalidString // в строке недопустимый символ, возврат с ошибкой
			}
		}

		rz[i].ch = simv
		if rz[i].type_simv == 'c' || rz[i].type_simv == 's' {
			rz[i].len = 1
		} // кол-во повторений (инициализация)

		// далее проверяем на 2 подряд цифры, в случае обнаружения возврат
		if i > 0 && rz[i].type_simv == 'd' && type_prev_simv == 'd' {
			return StrigConv.String(), ErrInvalidString
		}

		if i > 0 && rz[i].type_simv == 'd' && (rz[i-1].type_simv == 'c' || rz[i-1].type_simv == 's') {
			rz[i-1].len = (int)(rz[i].ch - '0')
			type_prev_simv = 'd'
		} else {
			type_prev_simv = rz[i].type_simv
			i++
		}

	}

	//	fmt.Printf("<")
	for j := 0; j < i; j++ { // формирование выходной строки
		for k := 0; k < rz[j].len; k++ {
			//			fmt.Printf("%c", rz[j].ch)
			StrigConv.WriteRune(rz[j].ch)
		}
	}

	//	fmt.Printf(">")
	return StrigConv.String(), nil
}
