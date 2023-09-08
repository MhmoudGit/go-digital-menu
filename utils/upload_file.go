package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request, fieldName string) string {
	// Get the file from the form
	file, _, err := r.FormFile(fieldName) // "file" is the name of the file input field in your form
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return ""
	}
	defer file.Close()

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
	return uniqueFilename
}

func generateUniqueFilename() string {
	return fmt.Sprintf("%d.jpg", time.Now().UnixNano())
}
