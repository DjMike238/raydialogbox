package main

import (
	"regexp"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	nameboxRect = rl.Rectangle{
		X:      100,
		Y:      300,
		Width:  125,
		Height: 30,
	}

	textboxRect = rl.Rectangle{
		X:      100,
		Y:      330,
		Width:  600,
		Height: 100,
	}

	blinkerSquare = rl.Rectangle{
		X:      680,
		Y:      410,
		Width:  10,
		Height: 10,
	}

	charPrintSpeed = time.Duration(75)

	audioRx = regexp.MustCompile(`\w`)
)

func main() {
	rl.InitWindow(800, 450, "RayDialogBox - POC by Dj_Mike238")
	rl.SetTargetFPS(60)

	// Load external files (audio and text)
	initTone()
	loadCharacters("data/characters.json")
	loadDialogue("data/dialogue.json")

	// Count characters, lines and check if text was completely printed
	currentChar := 0
	currentLine := 0
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

			// Draw name and text boxes
			if current.Mood != Idle {
				if current.Position != Hidden {
					nameboxRect.X = getNamePos(current.Position)
					rl.DrawRectangleRec(nameboxRect, rl.White)
				}

				rl.DrawRectangleLinesEx(textboxRect, 4, rl.White)

				// Also, print name in name box
				rl.DrawText(
					currentCharacter.Name,
					int32(nameboxRect.X)+(int32(nameboxRect.Width)/2)-(rl.MeasureText(currentCharacter.Name, 20)/2),
					int32(nameboxRect.Y)+5,
					20,
					rl.Black,
				)
			}

			// Prepare to print dialogue text
			txt := ""

			if !textDrawn && current.Mood != Idle {
				txt = current.Text[0 : currentChar+1]

				// Play blip tone on each valid character
				if audioRx.Match([]byte{current.Text[currentChar]}) {
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
				txt = current.Text

				// Check if blinker needs to be shown
				if current.Mood != Idle && !current.Autoplay {
					blinkStart <- 0
				}

				// Check if blinker needs to be drawn or not for the blinking effect
				if blinking && blinkNow {
					rl.DrawRectangleRec(blinkerSquare, rl.White)
				}

				// Check for pause on autoplay
				if current.Autoplay && current.Pause > 0 {
					time.Sleep(current.Pause * time.Millisecond)
				}

				// Reset vars for next line
				if rl.IsKeyPressed(rl.KeyEnter) || current.Autoplay {
					txt = ""
					textDrawn = false
					currentChar = 0
					currentLine += 1
					blinkStop <- 0
				}
			}

			// Print dialogue text in textbox
			rl.DrawText(
				txt,
				int32(textboxRect.X+10),
				int32(textboxRect.Y+10),
				20,
				rl.White,
			)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
