package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func drawBoxes(line DialogueLine) {
	if line.Mood != Idle {
		// Draw name and text boxes
		if line.NamePos != Hidden {
			setNameBoxPos(line.NamePos)
			rl.DrawRectangleRec(nameboxRect, rl.White)
		}

		rl.DrawRectangleLinesEx(textboxRect, 4, rl.White)

		// Print name in name box
		char := getCharacter(line.Name)
		charLen := rl.MeasureText(char.Name, TEXT_SIZE)

		rl.DrawText(
			char.Name,
			int32(nameboxRect.X)+(int32(nameboxRect.Width)/2)-(charLen/2),
			int32(nameboxRect.Y)+NAME_MARGIN_Y,
			TEXT_SIZE,
			rl.Black,
		)
	}
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
