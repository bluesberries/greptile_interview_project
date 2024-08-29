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

/**
 * This function was based on code from
 * https://docs.greptile.com/quickstart
 */
func query() {
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
	queryResp, err := queryClient.Do(queryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer queryResp.Body.Close()

	decodeQueryMessage(queryResp.Body)
}

type Source struct {
	Repository string
	Remote     string
	Branch     string
	FilePath   string
	LineStart  int
	LineEnd    int
	Summary    string
}

type QueryMessage struct {
	Sources []Source
}

/**
 * This function was based on code from
 * https://dev.to/taqkarim/you-might-not-be-using-json-decoder-correctly-in-golang-12mb
 */
func decodeQueryMessage(queryRespBody io.ReadCloser) {
	decoder := json.NewDecoder(queryRespBody)
	for {
		var m QueryMessage
		if err := decoder.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for i, source := range m.Sources {
			fmt.Printf("sources[%d]\n", i)
			fmt.Printf("repository:\t%s\n", source.Repository)
			fmt.Printf("remote:\t\t%s\n", source.Remote)
			fmt.Printf("branch:\t\t%s\n", source.Branch)
			fmt.Printf("filepath:\t%s\n", source.FilePath)
			fmt.Printf("linestart:\t%d\n", source.LineStart)
			fmt.Printf("lineend:\t%d\n", source.LineEnd)
			fmt.Printf("summary:\t%s\n\n", source.Summary)
		}
	}
}

func search() {
	searchData := map[string]interface{}{
		"messages": []map[string]string{
			{
				"id":      "some-id-1",
				"content": "Functions taht use recursion",
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

	decodeSearchMessage(searchResp.Body)
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

func decodeSearchMessage(searchRespBody io.ReadCloser) {
	decoder := json.NewDecoder(searchRespBody)
	for {
		var messages []SearchMessage
		if err := decoder.Decode(&messages); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

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
	}
}
