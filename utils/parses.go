package utils

import (
	"net/http"
	"strconv"
)

func ParseUint64(w http.ResponseWriter, str string) uint {
	uintStr, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		// Handle the error if the conversion fails
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return 0
	}
	return uint(uintStr)
}

func ParseMultipartForm(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusUnprocessableEntity)
		return
	}
}
