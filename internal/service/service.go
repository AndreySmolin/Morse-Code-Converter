package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectionMorse принимает и определяет тип  строки
// отдает преобразованную строку
func DetectionMorse(data string) (string, error) {
	sliceRune := []rune{',', ':', ';', '(', ')', '"', '/', '?', '#', '@', '=', '+', '_', '$', '¿', '¡', '&'}
	sliceData := strings.Split(data, "")
	if len(sliceData) == 0 {
		return "", errors.New("there is no data to display")
	}
	isMorse := false
	isText := false
	for _, s := range data {
		if 'a' <= s && s <= 'z' || 'A' <= s && s <= 'Z' {
			isText = true
		}
		if ('а' <= s && s <= 'я') || ('А' <= s && s <= 'Я') || s == 'ё' || s == 'Ё' {
			isText = true
		}
		if unicode.IsDigit(s) {
			isText = true
		}
		for _, x := range sliceRune {
			if s == x {
				isText = true
			}
		}
	}
	if isText {
		return morse.ToMorse(data), nil
	}
	for _, y := range data {
		if y == '.' || y == '-' {
			isMorse = true
		}
	}
	if isMorse {
		return morse.ToText(data), nil
	}
	return "", errors.New("unknown characters")
}
