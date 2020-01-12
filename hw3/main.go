package main

import (
	"fmt"
	"sort"
	"strings"
)

// Top10 returns top10 words in string
func Top10(input string) []string {
	// разбить строку
	words := strings.Fields(input)
	// пустой словарь для подсчета слов
	wordCountMap := map[string]int{}
	// подсчитываю кол-во слов
	for _, word := range words {
		_, ok := wordCountMap[word]
		if !ok {
			wordCountMap[word] = 1
		} else {
			wordCountMap[word] = wordCountMap[word] + 1
		}
	}
	// структура для слова
	type Word struct {
		Key   string
		Value int
	}
	// слайс для отсортированых слов
	var sorted []Word
	// наполнения слайса структурами слов
	for key, value := range wordCountMap {
		sorted = append(sorted, Word{key, value})
	}
	// сортировка слайса
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	// проверка длинны слайса чтобы возвращать топ10
	var sorted10 []Word
	if len(sorted) >= 10 {
		sorted10 = sorted[:10]
	} else {
		sorted10 = sorted
	}
	var top10 []string
	for _, word := range sorted10 {
		top10 = append(top10, word.Key)
	}
	return top10
}

func main() {
	top := Top10("one one one two two one two five six seven five ten")
	fmt.Println(top)
}
