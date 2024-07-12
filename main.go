package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var promptPath = "prompt.json"

func main() {
	URL := "https://api.openai.com/v1/chat/completions"

	loadEnv()
	apiKey := os.Getenv("API_KEY")

	prompt := readFile(promptPath)

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(prompt))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Println(string(body))

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func readFile(filePath string) []byte {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file. ", err)
	}
	return jsonData
}
