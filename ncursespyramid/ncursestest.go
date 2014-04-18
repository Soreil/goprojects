package main

import "code.google.com/p/goncurses"
import "log"

func main() {
	scr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()

	drawtriangle(scr, 20)
	scr.Refresh()
	scr.GetChar()
}

func drawtriangle(scr *goncurses.Window, height int) {
	for iter := 0; iter < height; iter++ {
		for runs := 1 + iter*2; runs > 0; runs-- {
			scr.MoveAddChar(iter, height-iter-2+runs, goncurses.ACS_DIAMOND)
		}
		scr.Refresh()
		goncurses.Nap(100 * iter)
	}
}
