package Task

// 1.

import (
	"strings"
)

//Алгоритм 1 задачи.
// 1) написать функцию, которая разбивает строку на слова
// 2) написать функцию, которая даляет все слова короче 3 символов
// 3)  написать функцию, которая возвращает map, где ключом является слово, а значением — количество

func TaskOne(s string) map[string]int {
	res := strings.Fields(s)
	message := make(map[string]int)

	for _, word := range res {
		if len(word) >= 3 {
			message[word]++
		}
	}

	return message

}

// Алгоритм 2 задачи
// 1) Реализовать фильтрацию с помощью условия
// 2) Преобразовать положительные цифры в квадрат
// 3) Реализовать сумму возведённых в квадрат цифр

func TaskTwo(numbers []int) int {
	filtered := []int{}
	for _, num := range numbers {
		if num >= 0 {
			filtered = append(filtered, num)
		}
	}

	squared := []int{}
	for _, num := range filtered {
		squared = append(squared, num*num)
	}

	sum := 0
	for _, num := range squared {
		sum += num
	}

	return sum
}

//Алгоритм для 3 задачи
// 1) инициализируем мапу в памяти
// 2) Скопировать элементы из первой мапы
// 3) Объединить элементы из второй мапы
// 4) Если ключ уже существует в result значения складываются
// 5) Если сумма больше 10, значение обновляется в result
// 6) Если сумма меньше или равна 10, ключ удаляется из result
// 7) Если ключ из второй мапы не существует в result, он добавляется только если его значение больше 10
// 8) Вернуть объединённую мапу

func TaskThree(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key, value := range map2 {
		if existingValue, ok := result[key]; ok {
			sum := existingValue + value
			if sum > 10 {
				result[key] = sum
			} else {
				delete(result, key)
			}
		} else {
			if value > 10 {
				result[key] = value
			}
		}
	}

	return result
}
