package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func main() {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", filesHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/clear-uploads", clearUploadsHandler)

	fmt.Println("🌿 Share Your Vegetables backend running on :8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func clearUploadsHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "Cannot read directory", http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		os.Remove(filepath.Join(uploadDir, f.Name()))
	}

	w.Write([]byte("Cleared"))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(50 << 20) // 50 MB max
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Cannot open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join(uploadDir, filepath.Base(fileHeader.Filename)))
		if err != nil {
			http.Error(w, "Cannot create file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		io.Copy(dst, file)
	}

	w.Write([]byte("Uploaded"))
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "Cannot read directory", http.StatusInternalServerError)
		return
	}

	type FileInfo struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
	}

	var list []FileInfo
	for _, f := range files {
		info, _ := f.Info()
		list = append(list, FileInfo{Name: f.Name(), Size: info.Size()})
	}

	json.NewEncoder(w).Encode(list)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	filename := filepath.Base(r.URL.Path)
	filepath := filepath.Join(uploadDir, filename)

	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, file)
}
