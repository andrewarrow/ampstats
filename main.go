package main

import (
	"ampstats/screen"
	"ampstats/util"
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
		t := screen.NewTemplate()
		t.Run()
	} else if command == "args" {
		fmt.Println("argMap", argMap)
		arg2 := util.GetArg(2)
		fmt.Println("arg2", arg2)
	} else if command == "help" {
		PrintHelp()
	}
}
