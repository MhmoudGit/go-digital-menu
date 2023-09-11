package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func ParseUint64(w http.ResponseWriter, str string) uint {
	uintStr, err := strconv.Atoi(str)
	if err != nil {
		// Handle the error if the conversion fails
		err := fmt.Sprintf("%v should be of type uint", str)
		http.Error(w, err, http.StatusUnprocessableEntity)
		return 0
	}
	return uint(uintStr)
}

func Parseint(w http.ResponseWriter, str string) int {
	intStr, err := strconv.Atoi(str)
	if err != nil {
		// Handle the error if the conversion fails
		err := fmt.Sprintf("%v should be of type int", str)
		http.Error(w, err, http.StatusUnprocessableEntity)
		return 0
	}
	return intStr
}

func ParseMultipartForm(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusUnprocessableEntity)
		return
	}
}

func ParseTime(dateString string) time.Time{
	// Parse the string into a time.Time value
	parsedTime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return parsedTime
	}
	return parsedTime
}
