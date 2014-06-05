package main

import (
	"code.google.com/p/goncurses"
	"fmt"
	"log"
)

var maxy, maxx int
var y, x int

func game(scr *goncurses.Window) {
	setup(scr)
	maxy, maxy = scr.MaxYX()
	for {
		move := input(scr)
		y, x := scr.CursorYX()
		switch move {
		case goncurses.KEY_LEFT:
			if x-1 > 0 {
				x -= 1
			}
		case goncurses.KEY_UP:
			if y+1 >= maxy {
				y += 1
			}
		case goncurses.KEY_DOWN:
			if y-1 > 0 {
				y -= 1
			}
		case goncurses.KEY_RIGHT:
			if x+1 >= maxx {
				x += 1
			}
		default:
			goncurses.End()
			fmt.Println(move, " ", goncurses.KEY_UP)
			log.Fatal()
		}
		scr.MoveAddChar(y, x, goncurses.ACS_DIAMOND)
		scr.Refresh()
	}
}

func setup(scr *goncurses.Window) {
	y, x = scr.MaxYX()
	scr.MoveAddChar(y/2, x/2, goncurses.ACS_DIAMOND)
	scr.Timeout(-1)
	// puts the cursor in the center
}

func input(scr *goncurses.Window) goncurses.Key {
	return scr.GetChar()
}

func main() {
	scr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	goncurses.Echo(false)
	scr.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	// Draws the outline
	scr.Refresh()
	game(scr)
	scr.GetChar()
	defer goncurses.End()
}
