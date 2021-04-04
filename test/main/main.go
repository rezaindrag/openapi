package main

import (
	"context"
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

	for _, f := range feed {
		fmt.Println(f.Article.GetTitle(), f.Person.GetName())
	}

	feedHome, resp, err := openapiClient.FeedHomeApi.FetchFeedHome(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Request.URL.String())

	for _, f := range feedHome {
		fmt.Println(f.Article.GetTitle(), f.PersonHome.GetFirstName(), f.PersonHome.GetLastName())
	}
}
