package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
	Поиск анаграмм по словарю
	Написать функцию поиска всех множеств анаграмм по словарю.

	Например:
	'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
	'листок', 'слиток' и 'столик' - другому.

	Требования:
	- Входные данные для функции: ссылка на массив, каждый элемент
	которого - слово на русском языке в кодировке utf8
	- Выходные данные: ссылка на мапу множеств анаграмм
	- Ключ - первое встретившееся в словаре слово из множества.
	- Значение - ссылка на массив, каждый элемент которого, слово из множества
	- Массив должен быть отсортирован по возрастанию.
	- Множества из одного элемента не должны попасть в результат.
	- Все слова должны быть приведены к нижнему регистру.
	- В результате каждое слово должно встречаться только один раз.
*/

// FindAnagram функция которая ищет анаграммы
func FindAnagram(arr []string) *map[string][]string {
	result := make(map[string][]string)

	for _, w := range arr {
		// приводим к нижнему регистру
		w := strings.ToLower(w)
		// преобразуем в рукны и сортируем
		word := []rune(w)
		unique := true
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		// проверем есть ли для этого слова подмножество
		for key := range result {
			tmp := []rune(key)
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] < tmp[j]
			})
			if string(tmp) == string(word) {
				result[key] = append(result[key], w)
				unique = false
				break
			}
		}
		// если нет создаем
		if unique {
			result[w] = []string{w}
		}
	}
	// удаляем дубликаты внути массива и множества из одного элемента
	for key := range result {
		result[key] = DelDuplicate(result[key])
		if len(result[key]) == 1 {
			delete(result, key)
		}
	}
	// сортируем массивы
	for _, v := range result {
		sort.Strings(v)
	}
	return &result
}

// DelDuplicate удаляет дубликаты
func DelDuplicate(s1 []string) []string {
	m := make(map[string]struct{})
	var res []string

	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

func main() {
	dict := []string{"пятак", "пятка", "тяпка", "листок", "столик", "слиток", "слово", "слово", "пятка"}
	fmt.Println(FindAnagram(dict))
}
