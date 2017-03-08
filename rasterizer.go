package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	frameRate = 60
)

// Find the smallest and greatest of two given numbers
func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// Draw a line
func Line(renderer *sdl.Renderer, x1, y1, x2, y2 int) {
	if x1 == x2 || y1 == y2 {
		return
	}

	startx, stopx := minmax(x1, x2)
	starty, stopy := minmax(y1, y2)

	xdiff := stopx - startx
	ydiff := stopy - starty

	if xdiff > ydiff {
		// We're going along X
		y := float32(starty)
		ystep := float32(ydiff) / float32(xdiff)
		if y1 != starty {
			// Move in the other direction along Y
			ystep = -ystep
			y = float32(stopy)
		}
		// Draw the line
		for x := startx; x < stopx; x++ {
			renderer.DrawPoint(x, int(y))
			y += ystep
		}
	} else {
		// We're going along Y
		x := float32(startx)
		xstep := float32(xdiff) / float32(ydiff)
		if x1 != startx {
			// Move in the other direction along X
			xstep = -xstep
			x = float32(stopx)
		}
		// Draw the line
		for y := starty; y < stopy; y++ {
			renderer.DrawPoint(int(x), y)
			x += xstep
		}
	}
}

// Draw a line where the coordinates are doubled and the pixels too
func DoublePixelLine(renderer *sdl.Renderer, x1, y1, x2, y2 int) {
	if x1 == x2 || y1 == y2 {
		return
	}

	startx, stopx := minmax(x1, x2)
	starty, stopy := minmax(y1, y2)

	xdiff := stopx - startx
	ydiff := stopy - starty

	doublex := 0
	doubley := 0

	if xdiff > ydiff {
		// We're going along X
		y := float32(starty)
		ystep := float32(ydiff) / float32(xdiff)
		if y1 != starty {
			// Move in the other direction along Y
			ystep = -ystep
			y = float32(stopy)
		}
		// Draw the line
		for x := startx; x < stopx; x++ {
			doublex = x * 2
			doubley = int(y) * 2
			renderer.DrawPoint(doublex, doubley)
			renderer.DrawPoint(doublex+1, doubley)
			renderer.DrawPoint(doublex, doubley+1)
			renderer.DrawPoint(doublex+1, doubley+1)
			y += ystep
		}
	} else {
		// We're going along Y
		x := float32(startx)
		xstep := float32(xdiff) / float32(ydiff)
		if x1 != startx {
			// Move in the other direction along X
			xstep = -xstep
			x = float32(stopx)
		}
		// Draw the line
		for y := starty; y < stopy; y++ {
			doublex = int(x) * 2
			doubley = y * 2
			renderer.DrawPoint(doublex, doubley)
			renderer.DrawPoint(doublex+1, doubley)
			renderer.DrawPoint(doublex, doubley+1)
			renderer.DrawPoint(doublex+1, doubley+1)
			x += xstep
		}
	}
}

// Checks if ESC is pressed
func escPressed() bool {
	keyMap := sdl.GetKeyboardState()
	return 1 == keyMap[sdl.K_ESCAPE]
}

func run() int {

	sdl.Init(sdl.INIT_VIDEO)

	// The actual resolution (screenspace)
	W := 800
	H := 600

	// The virtual resolution (worldspace)
	MAXX := 400
	MAXY := 300

	var (
		window   *sdl.Window
		renderer *sdl.Renderer
		err      error
	)

	window, err = sdl.CreateWindow("Random Pixelated Lines", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		window.Destroy()
	}()

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		renderer.Destroy()
	}()

	renderer.Clear()
	renderer.SetDrawColor(255, 0, 0, 255)

	rand.Seed(time.Now().UnixNano())

	var event sdl.Event
	running := true
	for running {
		// Innerloop
		renderer.SetDrawColor(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255)
		DoublePixelLine(renderer, rand.Intn(MAXX), rand.Intn(MAXY), rand.Intn(MAXX), rand.Intn(MAXY))
		renderer.Present()
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		sdl.Delay(1000 / frameRate)
	}

	return 0
}

func main() {
	os.Exit(run())
}
