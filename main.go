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

func main() {
	URL := "https://api.openai.com/v1/chat/completions"

	loadEnv()
	apiKey := os.Getenv("API_KEY")
	fmt.Println("API Key: ", apiKey)

	jsonData := []byte(`{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "system",
				"content": "You are a poetic assistant, skilled in explaining complex programming concepts with creative flair."
			},
			{
				"role": "user",
				"content": "Compose a poem that explains the concept of recursion in programming."
			}
		]
	}`)

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
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
