package services

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// fetch videos from youtube API
func FetchSvc(query string, maxResults int) map[string]string {
	apiKey := os.Getenv("API_KEY")

	// setting up youtube service
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("An error occured while setting up api client: %v", err)
	}

	// setting up call
	call := service.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(int64(maxResults))
	response, err := call.Do()
	if err != nil {
		log.Fatalf("An error occured while fetching videos: %v", err)
	}

	// filtering out videos from result
	videos := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	return videos
}
