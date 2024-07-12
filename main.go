package main

import (
	"bytes"
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

	fmt.Println(string(body))

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
