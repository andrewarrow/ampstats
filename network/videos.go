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

// API keys are not supported by this API
// TODO use oauth header
func DownloadCaptions(id string) string {

	key := os.Getenv("YOUTUBE_OAUTH")
	json := DoGetWithBearer(fmt.Sprintf("captions/%s", id), key)
	return json
}
