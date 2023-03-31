package main

import (
	"encoding/json"
	"os"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	characters []Character
	dialogue   []DialogueLine

	blip, medBlip, highBlip rl.Sound
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
