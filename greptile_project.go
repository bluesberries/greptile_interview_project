package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	greptileApiKey := os.Getenv("GREPTILE_API_KEY")
	githubToken := os.Getenv("GITHUB_TOKEN")

	queryData := map[string]interface{}{
		"messages": []map[string]string{
			{
				"id":      "some-id-1",
				"content": "Where's the entry point?",
				"role":    "user",
			},
		},
		"repositories": []map[string]string{
			{
				"repository": "bluesberries/redis-greptile-assessment",
				"branch":     "unstable",
			},
		},
		"sessionId": "test-session-id", // optional
	}
	queryPayload, _ := json.Marshal(queryData)

	queryReq, _ := http.NewRequest("POST", "https://api.greptile.com/v2/query", bytes.NewBuffer(queryPayload))
	queryReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", greptileApiKey))
	queryReq.Header.Set("X-Github-Token", githubToken)
	queryReq.Header.Set("Content-Type", "application/json")

	queryClient := &http.Client{}
	queryResp, _ := queryClient.Do(queryReq)
	defer queryResp.Body.Close()

	body, _ := io.ReadAll(queryResp.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
