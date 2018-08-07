// Example of software rendering
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// Resolution
	W = 1024
	H = 768

	// Size of pixels
	SCALE = 8

	// Target framerate
	frameRate = 60
)

// Find the smallest and greatest of two given numbers
func minmax(a, b int32) (int32, int32) {
	if a < b {
		return a, b
	}
	return b, a
}

// Draw a line
func Line(renderer *sdl.Renderer, x1, y1, x2, y2 int32) {
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
			renderer.DrawPoint(int32(x), int32(y))
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
			renderer.DrawPoint(int32(x), int32(y))
			x += xstep
		}
	}
}

// Draw a line where the coordinates are doubled and the pixels too
func ScaledPixelLine(renderer *sdl.Renderer, x1, y1, x2, y2 int32, scale int32) {
	if x1 == x2 || y1 == y2 {
		return
	}

	startx, stopx := minmax(x1, x2)
	starty, stopy := minmax(y1, y2)

	xdiff := stopx - startx
	ydiff := stopy - starty

	var doublex int32
	var doubley int32

	if xdiff > ydiff {
		// We're going along X
		ystep := float32(ydiff) / float32(xdiff)
		y := float32(starty)
		if y1 != starty {
			// Move in the other direction along Y
			ystep = -ystep
			y = float32(stopy)
		}
		// Draw the line
		for x := startx; x < stopx; x++ {
			doublex = x * scale
			doubley = int32(y) * scale
			renderer.FillRect(&sdl.Rect{doublex, doubley, scale, scale})
			y += ystep
		}
	} else {
		// We're going along Y
		xstep := float32(xdiff) / float32(ydiff)
		x := float32(startx)
		if x1 != startx {
			// Move in the other direction along X
			xstep = -xstep
			x = float32(stopx)
		}
		// Draw the line
		for y := starty; y < stopy; y++ {
			doublex = int32(x) * scale
			doubley = y * scale
			renderer.FillRect(&sdl.Rect{doublex, doubley, scale, scale})
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

	// The virtual resolution (worldspace)
	MAXX := W / SCALE
	MAXY := H / SCALE

	var (
		window   *sdl.Window
		renderer *sdl.Renderer
		err      error
	)

	window, err = sdl.CreateWindow("Random Pixelated Lines", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(W), int32(H), 0)
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
		ScaledPixelLine(renderer, int32(rand.Intn(MAXX)), int32(rand.Intn(MAXY)), int32(rand.Intn(MAXX)), int32(rand.Intn(MAXY)), int32(SCALE))
		renderer.Present()
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				ke := event.(*sdl.KeyboardEvent)
				if ke.Type == sdl.KEYDOWN {
					ks := ke.Keysym
					switch ks.Sym {
					case sdl.K_ESCAPE:
						running = false
					case sdl.K_q:
						running = false
					}
				}
			}
		}
		sdl.Delay(1000 / frameRate)
	}

	return 0
}

func main() {
	os.Exit(run())
}
