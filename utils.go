package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/raylib-go/raylib"
)

func check(fn string, err error) {
	if err != nil {
		fmt.Printf("error in %s: %s\n", fn, err)
		os.Exit(1)
	}
}

func cutText(text string, lines int) (cut string) {
	if strings.Count(text, "\n") > 2 {
		split := strings.Split(text, "\n")
		return strings.Join(split[lines-2:], "\n")
	}

	return text
}

func drawText(text string) {
	rl.DrawText(
		text,
		int32(textboxRect.X)+TEXT_MARGIN,
		int32(textboxRect.Y)+TEXT_MARGIN,
		TEXT_SIZE,
		rl.White,
	)
}

func isNextPressed() bool {
	return rl.IsKeyPressed(rl.KeyEnter) || rl.IsMouseButtonPressed(rl.MouseLeftButton)
}

func setNameBoxPos(pos NamePos) {
	switch pos {
	case Left:
		nameboxRect.X = 100
	case Center:
		nameboxRect.X = 337.5
	case Right:
		nameboxRect.X = 575
	}
}

func setPos(pos TextPos) {
	switch pos {
	case Top:
		nameboxRect.Y = NAMEBOX_TOP
		textboxRect.Y = TEXTBOX_TOP
		blinkerSquare.Y = BLINKER_TOP

	default: // Bottom
		nameboxRect.Y = NAMEBOX_BOTTOM
		textboxRect.Y = TEXTBOX_BOTTOM
		blinkerSquare.Y = BLINKER_BOTTOM
	}
}
