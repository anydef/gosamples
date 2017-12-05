package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

const backgroundColor = termbox.ColorBlue
const instructionsColor = termbox.ColorYellow

// Layout
const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const boardWidth = 20
const boardHeight = 20
const cellWidth = 2
const boardEndX = boardStartX + boardWidth*cellWidth
const boardEndY = boardStartY + boardHeight
const instructionsStartX = boardEndX + defaultMarginWidth
const instructionsStartY = boardStartY

var instructions = []string{
	"Goal: Fill in 5 lines!",
	"",
	"left   Left",
	"right  Right",
	"up     Rotate",
	"down   Down",
	"space  Fall",
	"s      Start",
	"p      Pause",
	"esc,q  Exit",
	"",
	"Level: %v",
	"Lines: %v",
	"",
	"GAME OVER!",
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	termbox.Clear(backgroundColor, backgroundColor)
	render()
	termbox.Flush()
	var yAxis int = titleStartY

	for {
		select {
		case event := <-eventQueue:
			if event.Type == termbox.EventKey {
				switch {
				case event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD:
					return
				case event.Ch == 'p':
					tbprint(titleStartX, yAxis, instructionsColor, backgroundColor, "Pressed P")
					yAxis++
					termbox.Flush()
				case event.Ch == 'c':
					termbox.Clear(backgroundColor, backgroundColor)
					termbox.Flush()
				default:
					termbox.Flush()
					//fmt.Printf(string(event.Key))
					//fmt.Printf(string(event.Ch))
				}
			}
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func render() {
	termbox.Clear(backgroundColor, backgroundColor)
	width, height := termbox.Size()
	tbprint(0, 0, instructionsColor, backgroundColor, "Termbox ui sample")

	//termbox.SetCell(1, 1, '.', termbox.ColorBlack, termbox.ColorBlack)
	//termbox.SetCell(2, 1, '.', termbox.ColorBlack, termbox.ColorBlack)
	//
	//termbox.SetCell(3, 1, '.', termbox.ColorRed, termbox.ColorRed)
	//termbox.SetCell(4, 1, '.', termbox.ColorRed, termbox.ColorRed)

	var cellColor termbox.Attribute
	for y := 0; y < height; y++ {
		for x := 0; x < width; x = x + 2 {
			for i := 0; i < cellWidth; i++ {
				if ( x+i+y )%2 == 0 {
					cellColor = termbox.ColorBlack
				} else {
					cellColor = termbox.ColorRed
				}
				termbox.SetCell(x+i, y, '.', cellColor, cellColor)
			}
		}
	}
	termbox.Flush()

}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
