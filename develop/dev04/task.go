package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Anagrams(strs []string) map[string][]string {
	var ans = make(map[string][]string)
	var ind = make(map[string]string)

	for _, s := range strs {
		s = strings.ToLower(s)
		arr := []rune(s)

		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})

		if _, ok := ind[string(arr)]; !ok {
			ans[s] = append(ans[s], s)
			ind[string(arr)] = s
		} else {
			i := ind[string(arr)]
			ans[i] = append(ans[i], s)
		}
	}
	for i := range ans {
		if len(ans[i]) < 2 {
			delete(ans, i)
		} else {
			sort.Strings(ans[i])
		}
	}

	return ans
}

func main() {
	input := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик"}

	fmt.Println(Anagrams(input))
}
