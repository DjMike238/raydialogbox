package main

import (
	"time"
)

// Hypotetical moods for characters
type Mood string

const (
	Normal Mood = "normal"
	Idle        = "idle"
)

// Name box position
type NamePos string

const (
	Hidden NamePos = "hidden"
	Left           = "left"
	Center         = "center"
	Right          = "right"
)

// Blip tone to play when printing text
type Tone string

const (
	Silent Tone = "silent"
	Blip        = "blip"
	Med         = "med"
	High        = "high"
)

// Textbox position
type TextPos string

const (
	Bottom TextPos = "bottom"
	Top            = "top"
)

// Data type containing dialogue line data
type DialogueLine struct {
	NamePos  NamePos       `json:"name_position,omitempty"`
	TextPos  TextPos       `json:"text_position,omitempty"`
	Mood     Mood          `json:"mood,omitempty"`
	Name     string        `json:"character,omitempty"`
	Text     string        `json:"text,omitempty"`
	Pause    time.Duration `json:"pause,omitempty"`
	Autoplay bool          `json:"autoplay,omitempty"`
}

// Data type containing character data
type Character struct {
	Name string `json:"name"`
	Tone Tone   `json:"tone"`
}
