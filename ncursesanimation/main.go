package main

import (
	"code.google.com/p/goncurses"
	"log"
)

type Grid struct {
	Screen *goncurses.Window
	X      int
	Y      int
}

func main() {
	src, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	g := Grid{src, 24, 80}
	g.Screen.Refresh()
	g.Screen.GetChar()
	defer goncurses.End()
}
