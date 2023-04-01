package main

import (
	"regexp"
	"strings"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	nameboxRect = rl.Rectangle{
		X:      100,
		Width:  125,
		Height: 30,
	}

	textboxRect = rl.Rectangle{
		X:      100,
		Width:  600,
		Height: 98,
	}

	blinkerSquare = rl.Rectangle{
		X:      680,
		Width:  10,
		Height: 10,
	}

	letterRx = regexp.MustCompile(`\w`)

	charPrintSpeed = time.Duration(PRINT_SPEED)
)

const (
	PRINT_SPEED = 75

	NAME_MARGIN_Y = 5

	TEXT_MARGIN = 10
	TEXT_SIZE   = 20

	// Textbox Y values
	TEXTBOX_TOP    = 24
	TEXTBOX_BOTTOM = 330

	// Namebox Y values
	NAMEBOX_TOP    = 122
	NAMEBOX_BOTTOM = 300

	// Blinker Y values
	BLINKER_TOP    = 102
	BLINKER_BOTTOM = 408
)

func main() {
	rl.InitWindow(800, 450, "RayDialogue - POC by Dj_Mike238")
	rl.SetTargetFPS(60)

	// Load external files (audio and text)
	initTone()
	loadCharacters("data/characters.json")
	loadDialogue("data/dialogue.json")

	// Count characters, lines and check if text was completely printed
	currentChar := 0
	currentLine := 0
	linesDrawn := 0
	textDrawn := false

	// Init channels for blinker
	blinkStart := make(chan uint8)
	blinkStop := make(chan uint8)

	// Start blinker checker
	blinking := false
	blinkNow := false

	go func() {
		tick := 400 * time.Millisecond
		blinkTick := time.Tick(tick)

		for {
			select {
			case <-blinkStart:
				blinking = true

			case <-blinkTick:
				blinkNow = !blinkNow

			case <-blinkStop:
				blinking = false
			}
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if currentLine < len(dialogue) {
			current := dialogue[currentLine]
			currentCharacter := getCharacter(current.Name)
			setPos(current.TextPos)

			if current.Mood != Idle {
				// Draw name and text boxes
				if current.NamePos != Hidden {
					nameboxRect.X = getNameBoxPos(current.NamePos)
					rl.DrawRectangleRec(nameboxRect, rl.White)
				}

				rl.DrawRectangleLinesEx(textboxRect, 4, rl.White)

				// Also, print name in name box
				characterLen := rl.MeasureText(currentCharacter.Name, TEXT_SIZE)

				rl.DrawText(
					currentCharacter.Name,
					int32(nameboxRect.X)+(int32(nameboxRect.Width)/2)-(characterLen/2),
					int32(nameboxRect.Y)+NAME_MARGIN_Y,
					TEXT_SIZE,
					rl.Black,
				)
			}

			if !textDrawn && current.Mood != Idle {
				// Check if the text is longer than 3 lines
				cut := cutText(current.Text[0:currentChar+1], linesDrawn)

				// Count lines drawn
				if current.Text[currentChar] == '\n' {
					linesDrawn += 1
				}

				// Print dialogue text in textbox
				drawText(cut)
				rl.EndDrawing()

				// Play blip tone on each valid character
				if letterRx.Match([]byte{current.Text[currentChar]}) {
					playTone(currentCharacter.Tone)
				}

				// Wait <charPrintSpeed> milliseconds before printing text
				time.Sleep(charPrintSpeed * time.Millisecond)

				if currentChar == len(current.Text)-1 {
					textDrawn = true
				} else {
					currentChar += 1
				}

			} else {
				// Print dialogue text in textbox
				cut := cutText(current.Text, linesDrawn)
				drawText(cut)

				// Check if blinker needs to be shown
				if current.Mood != Idle && !current.Autoplay {
					blinkStart <- 0
				}

				// Check if blinker needs to be drawn or not for the blinking effect
				if blinking && blinkNow {
					rl.DrawRectangleRec(blinkerSquare, rl.White)
				}

				rl.EndDrawing()

				// Check for pause on autoplay
				if current.Autoplay && current.Pause > 0 {
					time.Sleep(current.Pause * time.Millisecond)
				}

				// Reset vars for next line
				if isNextPressed() || current.Autoplay {
					textDrawn = false
					linesDrawn = 0
					currentChar = 0
					currentLine += 1
					blinkStop <- 0
				}
			}
		} else {
			rl.EndDrawing()
		}
	}

	rl.CloseWindow()
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

func getNameBoxPos(pos NamePos) float32 {
	switch pos {
	case Left:
		return 100
	case Center:
		return 337.5
	case Right:
		return 575
	}

	return 0
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
