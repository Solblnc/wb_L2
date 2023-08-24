package main

import (
	"fmt"
	"sort"
	"strings"
)

//Написать функцию поиска всех множеств анаграмм по словарю.
//
//
//Например:
//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.
//
//
//Требования:
//Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
//Выходные данные: ссылка на мапу множеств анаграмм
//Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//слово из множества.
//Массив должен быть отсортирован по возрастанию.
//Множества из одного элемента не должны попасть в результат.
//Все слова должны быть приведены к нижнему регистру.
//В результате каждое слово должно встречаться только один раз.

func set(words []string) map[string][]string {
	sortedSlice := make([]string, 0)

	for _, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedSlice = append(sortedSlice, string(runes))
	}

	setOfWords := make(map[string][]string)
	for index, value := range sortedSlice {
		setOfWords[value] = append(setOfWords[value], words[index])
	}
	resMaps := make(map[string]map[string]struct{})
	for _, value := range setOfWords {
		if len(value) == 0 {
			continue
		}
		resMaps[value[0]] = make(map[string]struct{})
		for _, str := range value {
			resMaps[value[0]][str] = struct{}{}
		}
	}
	resultSlice := make(map[string][]string)
	for key, value := range resMaps {
		for str := range value {
			resultSlice[key] = append(resultSlice[key], str)
		}
	}
	result := make(map[string][]string)
	for key, value := range resultSlice {
		sort.Slice(value, func(i, j int) bool {
			return strings.Compare(value[i], value[j]) == -1
		})
		result[key] = value
	}
	return result
}

func main() {
	words := []string{"тяпка", "пятак", "пятка", "клоун", "локун"}
	fmt.Println(set(words))

}
