package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("file") // "file" is the name of the file input field in your form
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename for the uploaded file
	uniqueFilename := generateUniqueFilename()

	// Create a new file to save the uploaded image
	filePath := filepath.Join("./uploads", uniqueFilename)
	newFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file data", http.StatusInternalServerError)
		return
	}

	// return the file name
	fmt.Fprintf(w, "File uploaded and saved as: %s", uniqueFilename)
}

func generateUniqueFilename() string {
	return fmt.Sprintf("%d.jpg", time.Now().UnixNano())
}
