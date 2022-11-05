package network

import (
	"fmt"
	"net/url"
	"os"
)

func SearchWord(q string) string {

	key := os.Getenv("YOUTUBE_KEY")
	url := fmt.Sprintf("search?part=snippet&order=date&maxResults=50&q=%s&key=%s", url.QueryEscape(q), key)
	json := DoGet(url)
	return json
}
