package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	defaultTimeout := time.Duration(int64(5)) * time.Second
	httpClient := &http.Client{
		Timeout: defaultTimeout,
	}

	openapiClient := openapi.NewAPIClient(&openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{
				URL: "http://localhost:3000",
			},
		},
		DefaultHeader: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPClient: httpClient,
	})

	feed, resp, err := openapiClient.FeedApi.FetchFeed(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Request.URL.String())
	bFeed, _ := json.Marshal(feed)
	fmt.Println(string(bFeed))

	for _, f := range feed {
		fmt.Println(f.Article.GetTitle(), f.Person.GetName())
	}
}
