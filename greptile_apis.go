package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var greptileApiKey = os.Getenv("GREPTILE_API_KEY")
var githubToken = os.Getenv("GITHUB_TOKEN")

func search(queryMessage string) []SearchMessage {
	searchData := map[string]interface{}{
		"messages": []map[string]string{
			{
				"id":      "some-id-1",
				"content": queryMessage,
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
	searchPayload, _ := json.Marshal(searchData)

	searchReq, _ := http.NewRequest("POST", "https://api.greptile.com/v2/search", bytes.NewBuffer(searchPayload))
	searchReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", greptileApiKey))
	searchReq.Header.Set("X-Github-Token", githubToken)
	searchReq.Header.Set("Content-Type", "application/json")

	searchClient := &http.Client{}
	searchResp, err := searchClient.Do(searchReq)
	if err != nil {
		log.Fatal(err)
	}
	defer searchResp.Body.Close()

	searchMessages := decodeSearchMessage(searchResp.Body)
	return searchMessages
}

type SearchMessage struct {
	Repository string
	Remote     string
	Branch     string
	FilePath   string
	LineStart  int
	LineEnd    int
	Summary    string
}

func decodeSearchMessage(searchRespBody io.ReadCloser) []SearchMessage {
	decoder := json.NewDecoder(searchRespBody)
	var allMessages []SearchMessage
	for {
		var messages []SearchMessage
		err := decoder.Decode(&messages)
		allMessages = append(allMessages, messages...)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		/*
			for i, m := range messages {
				fmt.Printf("[%d]\n", i)
				fmt.Printf("repository:\t%s\n", m.Repository)
				fmt.Printf("remote:\t\t%s\n", m.Remote)
				fmt.Printf("branch:\t\t%s\n", m.Branch)
				fmt.Printf("filepath:\t%s\n", m.FilePath)
				fmt.Printf("linestart:\t%d\n", m.LineStart)
				fmt.Printf("lineend:\t%d\n", m.LineEnd)
				fmt.Printf("summary:\t%s\n\n", m.Summary)
			}
		*/
	}
	return allMessages

}
