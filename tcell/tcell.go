package main

import (
	"github.com/gdamore/tcell"
	"fmt"
	"os"
	"time"
	"math/rand"
)

func main() {
	screen, err := tcell.NewScreen()

	defer screen.Fini()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite))
	screen.Clear()

	quit := make(chan struct{})

	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	cnt := 0
	dur := time.Duration(0)

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Microsecond * 500):
		}
		//screen.Clear()
		start := time.Now()
		draw(screen, cnt)
		cnt++
		dur += time.Now().Sub(start)
	}

	fmt.Printf("Finished %d boxes in %s\n", cnt, dur)
	fmt.Printf("Average is %0.3f ms / box\n", (float64(dur)/float64(cnt))/1000000.0)

}

func draw(screen tcell.Screen, count int) {
	width, height := screen.Size()

	if width == 0 || height == 0 {
		return
	}

	var y int = 0
	var x int = 0

	x = count % width
	y = count / width

	if y >= height *10 {
		return
	}

	style := tcell.StyleDefault

	rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
	style = style.Background(rgb)

	screen.SetCell(x, y, style, ' ')
	screen.Show()

}
