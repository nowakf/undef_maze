package main

//tiles
import (
	"fmt"
	"math/rand"

	"github.com/nsf/termbox-go"

	"time"
)

func main() {

	err := termbox.Init()
	defer termbox.Close()

	if err != nil {
		panic(err)
	}

	m := maze{}

	rand.Seed(time.Now().UnixNano())

	err = m.Convert(testmaze)

	if err != nil {
		panic(err)
	}

	fps := time.NewTicker(time.Second / 24)

	evchan := make(chan termbox.Event)

	termbox.SetInputMode(termbox.InputMouse | termbox.InputEsc)

	go func(evchan chan termbox.Event) {
		for {
			evchan <- termbox.PollEvent()
		}
	}(evchan)

	var mouseX, mouseY int

mainloop:
	for {
		m.Swap()

		m.Update()

		select {
		case event := <-evchan:
			if event.Type == termbox.EventKey {
				switch event.Key {
				case termbox.KeyCtrlC:
					break mainloop
				case termbox.KeyEsc:
					break mainloop
				}
			}
			if event.Type == termbox.EventMouse {
				mouseX, mouseY = event.MouseX, event.MouseY
				m.Query(mouseX, mouseY)
			}
		case <-fps.C:
			foodCounter(&m)
		}
		m.Draw()
		termbox.SetCell(mouseX, mouseY, '+', termbox.ColorBlack, termbox.ColorDefault)
		termbox.Flush()
		termbox.Clear(termbox.ColorBlack, termbox.ColorDefault)
	}
}

func foodCounter(m *maze) {
	foodLeft := fmt.Sprintf("food left: %d", len(m.food))
	w, h := termbox.Size()
	for i, letter := range foodLeft {
		termbox.SetCell(w-(len(foodLeft)-i), h-2, letter, termbox.ColorDefault, termbox.ColorBlack)
	}
}
