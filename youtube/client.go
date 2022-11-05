package youtube

import (
	"ampstats/network"
	"fmt"
)

type Video struct {
	Title string
}

func GetVideosByChannel(id string) []Video {
	items := []Video{}
	jsonString := network.VideosInChannel(id)
	fmt.Println(jsonString)
	return items
}
