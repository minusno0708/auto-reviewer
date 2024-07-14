package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var promptPath = "prompt.json"
var diffPath = "diff.txt"

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
	URL := "https://api.openai.com/v1/chat/completions"

	loadEnv()
	apiKey := os.Getenv("API_KEY")

	prompt := getPrompt()

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(prompt)))
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

	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		log.Fatal("Error unmarshaling JSON. ", err)
	}

	messageContext := gptResponse.Choices[0].Message.Content
	fmt.Println(messageContext)

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
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
