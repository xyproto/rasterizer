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
	// Size of "worldspace pixels", measured in "screenspace pixels"
	SCALE = 4

	// The resolution (worldspace)
	MAXX = 128
	MAXY = 128

	// Target framerate
	frameRate = 60

	// The resolution (screenspace)
	W = MAXX * SCALE
	H = MAXY * SCALE

	// Alpha value for opaque colors
	OPAQUE = 255
)

// isFullscreen checks if the current window has the WINDOW_FULLSCREEN_DESKTOP flag set
func isFullscreen(window *sdl.Window) bool {
	currentFlags := window.GetFlags()
	fullscreen1 := (currentFlags & sdl.WINDOW_FULLSCREEN_DESKTOP) != 0
	fullscreen2 := (currentFlags & sdl.WINDOW_FULLSCREEN) != 0
	return fullscreen1 || fullscreen2
}

// toggleFullscreen switches to fullscreen and back
// returns true if the mode has been switched to fullscreen
func toggleFullscreen(window *sdl.Window) bool {
	if !isFullscreen(window) {
		// Switch to fullscreen mode
		window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
		return true
	}
	// Switch to windowed mode
	window.SetFullscreen(sdl.WINDOW_SHOWN)
	return false
}

// toggleFullscreen2 switches to fullscreen and back, using a pointer to a boolean to keep track of the state
func toggleFullscreen2(window *sdl.Window, fullscreen *bool) {
	if *fullscreen {
		// Switch back to windowed mode
		window.SetFullscreen(sdl.WINDOW_SHOWN)
	} else {
		// Switch to fullscreen mode
		window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}
	*fullscreen = !*fullscreen
}

// ranb returns a random byte
func ranb() uint8 {
	return uint8(rand.Intn(255))
}

func run() int {

	sdl.Init(sdl.INIT_VIDEO)

	var (
		window    *sdl.Window
		renderer  *sdl.Renderer
		err       error
		ssOffsetX int32
		ssOffsetY int32
	)

	window, err = sdl.CreateWindow("Pixels!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(W), int32(H), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		window.Destroy()
	}()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		renderer.Destroy()
	}()

	renderer.SetDrawColor(0, 0, 0, OPAQUE)
	renderer.Clear()

	//surf, err := sdl.CreateRGBSurfaceWithFormat(0, MAXX, MAXY, 32, uint32(sdl.PIXELFORMAT_RGBA32))
	//if err != nil {
	//	panic(err)
	//}
	//surf.SetBlendMode(sdl.BLENDMODE_BLEND) // BLENDMODE_ADD or BLENDEMODE_MOD is also possible
	//surf.SetDrawColor(0, 0, 255, 127)
	//surf.Clear()

	rand.Seed(time.Now().UnixNano())

	var event sdl.Event
	running := true
	for running {
		// Innerloop
		renderer.SetDrawColor(ranb(), ranb(), ranb(), OPAQUE)
		ScaledPixelLine(renderer, rand.Int31n(MAXX), rand.Int31n(MAXY), rand.Int31n(MAXX), rand.Int31n(MAXY), int32(SCALE), ssOffsetX, ssOffsetY)
		renderer.SetDrawColor(ranb(), 0, 0, OPAQUE)
		ScaledPixel(renderer, rand.Int31n(MAXX), rand.Int31n(MAXY), int32(SCALE), ssOffsetX, ssOffsetY)
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
					case sdl.K_f, sdl.K_F11:
						if toggleFullscreen(window) {
							renderer.SetDrawColor(0, 0, 0, OPAQUE)
							renderer.Clear()
							dMode, err := window.GetDisplayMode()
							if err != nil {
								fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
								return 1
								panic(err)
							}
							// Screenspace offset, when in fullscreen mode
							ssOffsetX = (dMode.W - MAXX*SCALE) / 2
							ssOffsetY = (dMode.H - MAXY*SCALE) / 2
						} else {
							ssOffsetX = 0
							ssOffsetY = 0
						}
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
