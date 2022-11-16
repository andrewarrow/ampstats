package network

import (
	"fmt"
	"os"
)

func FetchCaptions(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("captions?key=%s&part=id,snippet&videoId=%s", key, id))
	return json
}

func DownloadCaption(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	json := DoGet(fmt.Sprintf("captions/%s?key=%s", id, key))
	return json
}
