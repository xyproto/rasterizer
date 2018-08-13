package main

import (
	"encoding/binary"
	"image/color"

	"github.com/xyproto/go-sdl2/sdl"
)

// Find the smallest and greatest of two given numbers
func minmax(a, b int32) (int32, int32) {
	if a < b {
		return a, b
	}
	return b, a
}

// Draw a line
func Line(pixels []uint32, x1, y1, x2, y2 int32, c color.RGBA, pitch int32) {
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
			Pixel(pixels, int32(x), int32(y), c, pitch)
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
			Pixel(pixels, int32(x), int32(y), c, pitch)
			x += xstep
		}
	}
}

// Draw a line where the coordinates are doubled and the pixels too
func ScaledPixelLine(renderer *sdl.Renderer, x1, y1, x2, y2, scale, ssOffsetX, ssOffsetY int32) {
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
			renderer.FillRect(&sdl.Rect{doublex + ssOffsetX, doubley + ssOffsetY, scale, scale})
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
			renderer.FillRect(&sdl.Rect{doublex + ssOffsetX, doubley + ssOffsetY, scale, scale})
			x += xstep
		}
	}
}

// Draw a worldspace pixel in screenspace
func ScaledPixel(renderer *sdl.Renderer, x, y, scale, ssOffsetX, ssOffsetY int32) {
	renderer.FillRect(&sdl.Rect{x*scale + ssOffsetX, y*scale + ssOffsetY, scale, scale})
}

// Draw a screenspace pixel
func Pixel(pixels []uint32, x, y int32, c color.RGBA, pitch int32) {
	pixels[y*pitch+x] = binary.BigEndian.Uint32([]uint8{c.A, c.R, c.G, c.B})
}
