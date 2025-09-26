package main

import (
	"log"
	"task-tracker/cmd"
)

func main() {
	root := cmd.RootCmd()

	if err := root.Execute(); err != nil {
		log.Println(err)
	}
}
