package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request, fieldName string) string {
	// Get the file from the form
	file, header, err := r.FormFile(fieldName) // "file" is the name of the file input field in your form
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return ""
	}
	defer file.Close()

	// Validate file type (allow only PNG and JPG)
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	ext := filepath.Ext(header.Filename)
	if !isValidExtension(ext, allowedExtensions) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return ""
	}

	// Generate a unique filename for the uploaded file
	uniqueFilename := generateUniqueFilename()

	// Create a new file to save the uploaded image
	filePath := filepath.Join("./uploads", uniqueFilename)
	newFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return ""
	}
	defer newFile.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file data", http.StatusInternalServerError)
		return ""
	}

	// return the file name
	return fmt.Sprintf("/uploads/%v", uniqueFilename)
}

func isValidExtension(extension string, allowedExtensions []string) bool {
	extension = strings.ToLower(extension)
	for _, ext := range allowedExtensions {
		if ext == extension {
			return true
		}
	}
	return false
}

func generateUniqueFilename() string {
	return fmt.Sprintf("%d.jpg", time.Now().UnixNano())
}
