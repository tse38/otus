package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func clearString(textIn string) string {
	// Очищаем текст от знаков препинания, переносов и табуляций
	var textOut strings.Builder // инициализация строки
	var chPrev rune
	for _, ch := range textIn {
		switch ch {
		case '.', ',', '!', '"', ':', '\n', '\t':
			ch = ' ' // замена знаков препинания на пробел
		}
		if !(ch == ' ' && chPrev == ' ') { // удаление лишних пробелов
			textOut.WriteRune(unicode.To(unicode.LowerCase, ch)) // заодно преобразуем в нижний регистр
		}
		chPrev = ch
	}
	//  получена строка, очищенная от знаков препинания, переводов строки, табуляций
	return textOut.String()
}

func Top10(textIn string, lenMin int) []string {
	var r []string
	if len(textIn) > 150000 {
		return r
	} // слишком большой текст
	if textIn == "" {
		return r
	} // пусто

	// Очищаем текст от знаков препинания, переносов и табуляций
	textOut := clearString(textIn)
	//  получена строка, очищенная от знаков препинания, переводов строки, табуляций
	//	fmt.Printf("\n%s",textOut.String())

	stroki := strings.Split(textOut, " ") // получаем массив слов

	// Делаем map  с ключём СЛОВО и значением "количество в тексте"
	// Тут же отбрасываем "-" как отдельный элемент (слово)
	mapSlovar := make(map[string]int, len(stroki)) // берем с запасом, расчитывая что все слова в тексте уникальные
	for _, slovo := range stroki {
		if !(slovo == "-" || slovo == "" || slovo == " " /*|| slovo == " —"*/) {
			mapSlovar[slovo]++
		}
	}

	// С сортировкай map-а по значению параметра проблемы, поэтому переписываем его в структуру.
	// Вариант сразу писать в нее без использования map-а есть, но реализация сложнее, и в данном случае смысла не имеет
	type tSortSlovar = struct {
		slovo string
		kolvo int
	}

	sortSlovar := make([]tSortSlovar, len(mapSlovar))

	iNew := 0
	for key, kolvo := range mapSlovar { // заполняем структуру из map-а для последующей сортировки
		sortSlovar[iNew].slovo = key
		sortSlovar[iNew].kolvo = kolvo
		iNew++
	}
	// далее сортируем стандартными средствами, сначала по количеству повторений, потом лексикографически
	sort.SliceStable(sortSlovar, func(i, j int) bool {
		if sortSlovar[i].kolvo == sortSlovar[j].kolvo { // здесь упорядочиваем по алфавиту
			return sortSlovar[i].slovo < sortSlovar[j].slovo
		}
		return sortSlovar[i].kolvo > sortSlovar[j].kolvo // а здесь по количеству повторений слова в тексте
	})

	stringTop10 := make([]string, 10)
	k := 0
	for _, sl := range sortSlovar { // формируем слайс top10
		if utf8.RuneCountInString(sl.slovo) >= lenMin { // проверяем на минимально учитываемую длину слова
			stringTop10[k] = sl.slovo
			//	    fmt.Printf("\n%d <%s> %d <%s>",k,sl.slovo, sl.kolvo, stringTop10[k])
			if k == 9 {
				break
			}
			k++
		}
	}

	return stringTop10
}
