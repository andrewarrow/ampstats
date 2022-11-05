package network

import (
	"ampstats/files"
	"fmt"
	"os"
)

func VideosInChannel(id string) string {

	cacheFile := fmt.Sprintf("cache/%s.json", id)
	cacheJson := files.ReadFile(cacheFile)
	if cacheJson != "" {
		fmt.Println("using cache")
		return cacheJson
	}
	key := os.Getenv("YOUTUBE_KEY")
	url := fmt.Sprintf("search?part=snippet&order=date&maxResults=50&key=%s&channelId=%s", key, id)
	json := DoGet(url)
	files.SaveFile(cacheFile, json)
	return json
}
