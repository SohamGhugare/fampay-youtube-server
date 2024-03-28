package services

import (
	"context"
	"log"
	"os"

	"github.com/SohamGhugare/fampay-youtube-server/models"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// fetch videos from youtube API
func FetchSvc(query string, maxResults int) []models.Video {
	apiKey := os.Getenv("API_KEY")

	// setting up youtube service
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("An error occured while setting up api client: %v", err)
	}

	// setting up call (query, maxResults and order by date)
	call := service.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(int64(maxResults)).Order("date")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("An error occured while fetching videos: %v", err)
	}

	// filtering out videos from result
	var videos []models.Video

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			video := models.Video{
				ID:           item.Id.VideoId,
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
				PublishedAt:  item.Snippet.PublishedAt,
			}
			videos = append(videos, video)
		}
	}

	return videos
}
