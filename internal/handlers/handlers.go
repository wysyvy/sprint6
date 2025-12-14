package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.ReadFile("index.html")

	if err != nil {
		http.Error(w, "Error load page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	w.Write(htmlFile)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error get file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error read file", http.StatusInternalServerError)
		return
	}
	input := string(data)

	res, err := service.Convert(input)
	if err != nil {
		http.Error(w, "Error convert", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(fileHeader.Filename)
	fileName := time.Now().UTC().String() + ext

	out, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error create res file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err = out.WriteString(res); err != nil {
		http.Error(w, "Error write res file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, res)
}
