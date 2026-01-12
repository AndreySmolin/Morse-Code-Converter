package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler обработчик метода GET возвращает ответ по шаблону html
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

// LoadHandler обработчик метода Post парсит форму html,
// обрабатывает полученный тип текста (Морзе или обычный),
// записывает в локальный файл и отправляет клиенту
func LoadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	line, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "reading error", http.StatusInternalServerError)
		return
	}
	str, err := service.DetectionMorse(string(line))
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}
	name := fmt.Sprintf("%s.txt", time.Now().UTC().String())
	if err := os.WriteFile(name, []byte(str), 0755); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(str))
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}
}
