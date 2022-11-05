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

	if command == "run" {
		t := screen.NewAmpStats()
		t.Run()
	} else if command == "videos" {
		for i, video := range youtube.GetVideosByChannel("AubreyMarcusPod") {
			fmt.Printf("%02d. %s\n", i, video.Title)
		}
	} else if command == "args" {
		fmt.Println("argMap", argMap)
		arg2 := util.GetArg(2)
		fmt.Println("arg2", arg2)
	} else if command == "help" {
		PrintHelp()
	}
}
