package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var URL = "https://api.openai.com/v1/chat/completions"

var promptPath = "prompt.json"
var diffPath = "diff.txt"

var savePath = "message.txt"

type GPTResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`

	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	loadEnv()
	apiKey := os.Getenv("API_KEY")

	prompt := getPrompt()

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(prompt)))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	respBody := sendRequest(req)
	message := extractMessage(respBody)

	saveMessage(message)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getPrompt() string {
	rawDiff := string(readFile(diffPath))
	basePrompt := string(readFile(promptPath))

	// " や \n などの特殊文字をエスケープ
	formatDiff := strconv.Quote(rawDiff)
	formatDiff = formatDiff[1 : len(formatDiff)-1]

	concatPrompt := strings.Replace(basePrompt, "[DIFF]", formatDiff, -1)

	return concatPrompt
}

func readFile(filePath string) []byte {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file. ", err)
	}
	return jsonData
}

func sendRequest(req *http.Request) []byte {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	return body
}

func extractMessage(body []byte) string {
	var gptResponse GPTResponse
	err := json.Unmarshal(body, &gptResponse)
	if err != nil {
		log.Fatal("Error unmarshaling JSON. ", err)
	}

	return gptResponse.Choices[0].Message.Content
}

func saveMessage(message string) {
	err := ioutil.WriteFile(savePath, []byte(message), 0644)
	if err != nil {
		log.Fatal("Error writing file. ", err)
	}
}
