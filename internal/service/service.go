package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectionMorse(data string) (string, error) {
	sliceData := strings.Split(data, "")
	if len(sliceData) == 0 {
		return "", errors.New("there is no data to display")
	}
	if sliceData[0] == "." || sliceData[0] == "-" {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}
