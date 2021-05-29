package scale

import (
	"strings"
)

var pitchesSharp = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var pitchesFlat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var notesRepresentation = map[string]string{
	"G":  "sharp",
	"D":  "sharp",
	"A":  "sharp",
	"E":  "sharp",
	"B":  "sharp",
	"b":  "sharp",
	"F#": "sharp",
	"f#": "sharp",
	"e":  "sharp",
	"G#": "sharp",
	"C#": "sharp",
	"A#": "flat",
	"D#": "flat",
	"F":  "flat",
	"g":  "flat",
	"Bb": "flat",
	"bb": "flat",
	"Eb": "flat",
	"d":  "flat",
	"Db": "flat",
}

//Scale generates the musical chromaticScale starting with the tonic and following the
//specified interval pattern
func Scale(tonic string, interval string) []string {
	scale := []string{strings.Title(tonic)}
	pitches := pitches(tonic)
	position := findPosition(tonic, pitches)
	if interval == "" {
		scale := pitches[position:]
		scale = append(scale, pitches[0:position]...)
		return scale
	}
	for _, step := range interval {
		switch step {
		case 'M':
			position += 2
		case 'm':
			position += 1
		case 'A':
			position += 3
		}
		scale = append(scale, pitches[position%12])
	}
	return scale[0 : len(scale)-1]
}

func pitches(tonic string) []string {
	if notesRepresentation[tonic] == ("flat") {
		return pitchesFlat
	}
	return pitchesSharp
}

func findPosition(tonic string, pitches []string) int {
	tonic = strings.Title(tonic)
	for i, pitch := range pitches {
		if pitch == tonic {
			return i
		}
	}
	return -1
}
