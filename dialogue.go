package main

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	characters []Character
	dialogue   []DialogueLine

	blip, medBlip, highBlip rl.Sound

	wrapRx = regexp.MustCompile(`\W+`)

	// Text margin * 3 = 1 from the left, 2 from the right for lower tolerance
	textLimit = int32(textboxRect.Width) - TEXT_MARGIN*3
)

func loadCharacters(file string) {
	cnt, err := os.ReadFile(file)
	check("loadCharacters", err)

	err = json.Unmarshal(cnt, &characters)
	check("loadCharacters", err)
}

func loadDialogue(file string) {
	cnt, err := os.ReadFile(file)
	check("loadDialogue", err)

	err = json.Unmarshal(cnt, &dialogue)
	check("loadDialogue", err)

	for i, v := range dialogue {
		v.Text = wrap(v.Text)
		dialogue[i] = v
	}
}

func getCharacter(name string) Character {
	for _, c := range characters {
		if c.Name == name {
			return c
		}
	}

	return Character{}
}

func getNamePos(pos NamePos) float32 {
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

func initTone() {
	rl.InitAudioDevice()

	blip = rl.LoadSound("snd/blip.wav")
	rl.SetSoundVolume(blip, 0.5)

	medBlip = rl.LoadSound("snd/med_blip.wav")
	rl.SetSoundVolume(medBlip, 0.5)

	highBlip = rl.LoadSound("snd/high_blip.wav")
	rl.SetSoundVolume(highBlip, 0.5)
}

func playTone(tone Tone) {
	switch tone {
	case Blip:
		rl.PlaySound(blip)
	case Med:
		rl.PlaySound(medBlip)
	case High:
		rl.PlaySound(highBlip)
	}
}

func wrap(text string) (wrapped string) {
	if rl.MeasureText(text, TEXT_SIZE) < textLimit {
		return text
	}

	words := wrapRx.Split(text, -1)
	symbols := wrapRx.FindAllString(text, -1)

	for i, word := range words {
		if rl.MeasureText(wrapped + word, TEXT_SIZE) > textLimit {
			wrapped += "\n"
		}

		wrapped += word

		if len(symbols) > i {
			if strings.HasSuffix(symbols[i], " ") {
				sym := symbols[i]
				if len(sym) > 1 {
					sym = strings.TrimSuffix(sym, " ")
				}

				length := rl.MeasureText(wrapped + sym, TEXT_SIZE)

				if length > textLimit {
					wrapped += sym + " \n"
				} else {
					wrapped += symbols[i]
				}
			}
		}
	}

	return wrapped + symbols[len(symbols)-1]
}
