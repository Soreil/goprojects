package main

import "github.com/nsf/termbox-go"
import "math"
import "math/rand"
import "strconv"
import "time"
import "unicode/utf8"

func draw(snakeX, snakeY *int, length int) {
	str := strconv.FormatInt(int64(length), 10)
	i := 0
	for length/int(math.Pow(float64(10), float64(i))) != 0 {
		r, _ := utf8.DecodeRuneInString(string(str[i]))
		termbox.SetCell(i, 0, r, termbox.ColorBlack, termbox.ColorWhite)
		i++
	}
	termbox.SetCell(*snakeX, *snakeY, '#', termbox.ColorBlack, termbox.ColorWhite)
	termbox.Flush()
}

func init_draw() (int, int, int, int) {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	termbox.SetCell(w/2, h/2, '#', termbox.ColorBlack, termbox.ColorWhite)
	return w / 2, h / 2, w, h
}

func spawn_item(w, h, length int) {
	rand.Seed(int64(length))
	randw := rand.Intn(w)
	randh := rand.Intn(h)
	termbox.SetCell(randw, randh, '*', termbox.ColorRed, termbox.ColorDefault)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	snakeX, snakeY, w, h := init_draw()
	direction := "right"
	length := 1
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowLeft {
				direction = "left"
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowDown {
				direction = "down"
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowUp {
				direction = "up"
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowRight {
				direction = "right"
			}
		default:
			switch direction {
			case "left":
				snakeX--
			case "down":
				snakeY++
			case "up":
				snakeY--
			case "right":
				snakeX++
			}

			if snakeX >= w || snakeY >= h || snakeX <= 0 || snakeY < 0 {
				snakeX, snakeY, w, h = init_draw()
				length = 1
			}
			length++
			spawn_item(w, h, length)
			draw(&snakeX, &snakeY, length)
			time.Sleep(time.Second / 60)
		}
	}
}
