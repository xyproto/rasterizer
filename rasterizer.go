package main

import (
	"math/rand"
	"time"
	//"fmt"	

	"github.com/banthar/Go-SDL/sdl"
)

// Find the smallest and greatest of two given numbers
func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// Draw a line
func Line(screen *sdl.Surface, x1 int, y1 int, x2 int, y2 int, color sdl.Color) {
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
			screen.Set(x, int(y), color)
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
			screen.Set(int(x), y, color)
			x += xstep
		}
	}
}

// Draw a line where the coordinates are doubled and the pixels too
func DoublePixelLine(screen *sdl.Surface, x1 int, y1 int, x2 int, y2 int, color sdl.Color) {
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
			screen.Set(doublex, doubley, color)
			screen.Set(doublex+1, doubley, color)
			screen.Set(doublex, doubley+1, color)
			screen.Set(doublex+1, doubley+1, color)
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
			screen.Set(doublex, doubley, color)
			screen.Set(doublex+1, doubley, color)
			screen.Set(doublex, doubley+1, color)
			screen.Set(doublex+1, doubley+1, color)
			x += xstep
		}
	}
}

// Checks if ESC is pressed
func escPressed() bool {
	keyMap := sdl.GetKeyState()
	return 1 == keyMap[sdl.K_ESCAPE]
}

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	// The actual resolution (screenspace)
	W := 800
	H := 600

	// The virtual resolution (worldspace)
	MAXX := 400
	MAXY := 300

	screen := sdl.SetVideoMode(W, H, 32, sdl.FULLSCREEN)
	if screen == nil {
		panic(sdl.GetError())
	}
	sdl.EnableUNICODE(1)
	sdl.ShowCursor(0)

	sdl.WM_SetCaption("Random Lines", "")

	red := sdl.Color{255, 0, 0, 255}
	color := red

	rand.Seed(time.Now().UnixNano())

	for {
		color = sdl.Color{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
		DoublePixelLine(screen, rand.Intn(MAXX), rand.Intn(MAXY), rand.Intn(MAXX), rand.Intn(MAXY), color)
		if escPressed() {
			break
		}
		screen.Flip()
		sdl.PollEvent()
		//sdl.Delay(10)
	}

	sdl.Quit()
}
