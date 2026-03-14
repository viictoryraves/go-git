package main

import (
	"fmt"
	"os"

	"gogit/data"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A command is needed")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		data.Init()
		wd, _ := os.Getwd()
		fmt.Printf("Initialized empty ugit repository in %s/%s", wd, data.GogitDir)
	case "hash-object":
		if len(os.Args) < 3 {
			panic("The path argument is missing")
		}
		data.HashObject(os.Args[2])
	}
}
