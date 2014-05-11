package main

import (
	"code.google.com/p/goncurses"
	"log"
)

func main() {
	src, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()
	src.Refresh()
	src.GetChar()
}
