package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ResponseData struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Amount    int    `json:"amount"`
	Refunded  int    `json:"refunded"`
	CreatedAt string `json:"createdAt"`
}

func GetPayment(id string) ResponseData {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	// Define the URL you want to make the GET request to.
	url := "https://api.moyasar.com/v1/payments/" + id

	// Create an HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ResponseData{}
	}

	// Set the Authorization header with the Bearer token
	token := os.Getenv("M_KEY")
	fmt.Println(token)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(token+":")))

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ResponseData{}
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.Status)
		return ResponseData{}
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ResponseData{}
	}

	// Unmarshal the JSON response into a struct
	var responseData ResponseData
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ResponseData{}
	}

	return responseData
}
