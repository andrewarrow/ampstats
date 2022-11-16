package main

import (
	"ampstats/screen"
	"ampstats/util"
	"ampstats/youtube"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap := util.ArgsToMap()
	os.Mkdir("cache", 0755)

	if command == "run" {
		t := screen.NewAmpStats()
		t.Run()
	} else if command == "search" {
		youtube.SearchWord("AubreyMarcusPod")
	} else if command == "videos" {
		for i, video := range youtube.GetVideosByChannel("UC604SM0YhltEKZ5hmDs_Gqw") {
			fmt.Printf("%02d. [%s] %s\n", i, video.Id, video.Title)
		}
	} else if command == "args" {
		fmt.Println("argMap", argMap)
		arg2 := util.GetArg(2)
		fmt.Println("arg2", arg2)
	} else if command == "help" {
		PrintHelp()
	}
}
