package Task

// 1.

import (
	"regexp"
	"sort"
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

// Алгоритм
//1 Удалить все знаки препинания из строки
//2 Перевести все буквы в нижний регистр
//3 Разбить строку на слова
//4 Подсчитать количество появлений каждого слова

func TaskFour(s string) map[string]int {

	reg := regexp.MustCompile(`[^\w\s]`)
	cleaned := reg.ReplaceAllString(s, "")
	lowered := strings.ToLower(cleaned)

	words := strings.Fields(lowered)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

// Алгоритм
//1 Создать новый слайс, где каждый элемент является суммой всех предыдущих элементов оригинального слайса
//2 Удалить все элементы меньше среднего арифметического этого слайса

func TaskFive(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// Накопление результата
	accumulated := make([]int, len(nums))
	accumulated[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		accumulated[i] = accumulated[i-1] + nums[i]
	}

	// Возвращаем только последний элемент, который является суммой всех элементов
	return []int{accumulated[len(accumulated)-1]}
}

// Алгоритм
//1 извлечь все уникальные символы из строки
//2 реверсировать строку

func TaskSix(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Алгоритм
//1 Разбить строку на слова
//2 Создать map, где ключ — слово, а значение — его длина
//3 Отсортировать map по длине слов и вернуть список ключей

func TaskSeven(s string) []string {
	words := strings.Split(s, " ")
	sort.Strings(words)
	return words
}
