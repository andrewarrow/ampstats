package youtube

import (
	"ampstats/network"
	"ampstats/parse"
	"fmt"
	"html"
)

type Video struct {
	Id    string
	Title string
}

func GetVideosByChannel(id string) []Video {
	items := []Video{}
	jsonString := network.VideosInChannel(id)
	//fmt.Println(jsonString)
	result := parse.ParseJson(jsonString)
	//fmt.Println(result.Items)
	for _, item := range result.Items {
		v := Video{}
		v.Title = html.UnescapeString(item.Snippet.Title)
		v.Id = item.Id.VideoId
		items = append(items, v)
	}
	return items
}

func SearchWord(word string) {
	jsonString := network.SearchWord(word)
	fmt.Println(jsonString)
}
