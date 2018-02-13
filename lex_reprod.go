package main

import (
	"strings"
	//"log"
	"errors"
	"math/rand"
	"time"
	"log"
	"strconv"
)

var (
	mat    = []string{"ху", "пизд"}
)

func huifma(word string) (string, error) {
	v_index := vowel_index(&word)
	if v_index != -1 {
		utf8word := []rune(strings.ToLower(word))
		part2 := utf8word[vowel_index(&word) : len(word)-1]
		switch part2[0] {
		case 'а':
			part2[0] = 'я'
		case 'о':
			part2[0] = 'ё'
		case 'э':
			part2[0] = 'е'
		case 'ы':
			part2[0] = 'и'
		case 'у':
			part2[0] = 'ю'
		}
		sr := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(sr)
		return mat[r1.Intn(len(mat))] + string(part2), nil
	}
	return strings.ToLower(word), errors.New("vowels not found")
}



func vowel_index(word *string) int {
	var i int
	for _, value := range *word {
		switch value {
		case 'а', 'о', 'и', 'е', 'ё', 'э', 'ы', 'у', 'ю', 'я', 'А', 'О', 'И', 'Е', 'Ё', 'Э', 'Ы', 'У', 'Ю', 'Я':
			return i
		}
		i++
	}
	return -1
}

func probably(per int) (bool) {
	if (per < 100) && (per > 0) {
		sr := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(sr)
		cur_n := r1.Intn(100)
		log.Println("rand "+ strconv.Itoa(cur_n) + " of " + strconv.Itoa(per))
		if cur_n < per {
			return true
		}
	}
	return false
}