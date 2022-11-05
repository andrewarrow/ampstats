package network

import (
	"fmt"
	"os"
)

func VideosInChannel(id string) string {

	key := os.Getenv("YOUTUBE_KEY")
	url := fmt.Sprintf("search?part=snippet&order=date&maxResults=50&key=%s&channelId=%s", key, id)
	json := DoGet(url)
	return json
}
