package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func drawText(text string) {
	rl.DrawText(
		text,
		int32(textboxRect.X)+TEXT_MARGIN,
		int32(textboxRect.Y)+TEXT_MARGIN,
		TEXT_SIZE,
		rl.White,
	)
}
