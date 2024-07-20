package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func semitoneInterval(input int) string {
	semiToInterval := map[int][]string{
		1:  {"Diminished second"},
		2:  {"Minor second"},
		3:  {"Major second", "Diminished third"},
		4:  {"Minor third", "Augmented second"},
		5:  {"Major third", "Diminished fourth"},
		6:  {"Perfect fourth", "Augmented third"},
		7:  {"Augmented fourth", "Diminished fifth"},
		8:  {"Perfect fifth", "Diminished sixth"},
		9:  {"Minor sixth", "Augmented fifth"},
		10: {"Major sixth", "Diminished seventh"},
		11: {"Minor seventh", "Augmented sixth"},
		12: {"Major seventh", "Diminished octave"},
		13: {"Perfect octave", "Augmented seventh"},
		14: {"Augmented octave", "Compound Diminished second"},
	}
	result := ""
	if input > 14 {
		input = input%13 + 1
		result += "Compound "
	}
	// Retrieve the interval descriptions for the given semitone count
	intervals, exists := semiToInterval[input]
	if !exists {
		return "Invalid semitone count"
	}

	// Concatenate all interval descriptions into a single string

	for i, interval := range intervals {
		if i > 0 {
			result += " or " // Add "or" to separate multiple possibilities
		}
		result += interval
	}

	return result
}
func Interval_main(noteinput []string) string {
	// List of musical notes
	notes_list := []string{
		"C", "C#\nDb", "D", "D#\nEb", "E", "F", "F#\nGb", "G",
		"G#\nAb", "A", "A#\nBb", "B", "C'", "C#\nDb'", "D'", "D#\nEb'",
		"E'", "F'", "F#\nGb'", "G'", "G#\nAb'", "A'", "A#\nBb'", "B'",
	}

	// Initialize slice to store indices
	indices := make([]int, 0)

	// Find indices of input notes in the notes_list
	for _, inputNote := range noteinput {
		found := false
		for idx, note := range notes_list {
			if note == inputNote {
				indices = append(indices, idx)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Note %s not found in the list\n", inputNote)
		}
	}

	// Calculate and print the semitone distance if there are enough indices
	if len(indices) == 2 {
		semi_distance := abs(indices[0]-indices[1]) + 1
		output := semitoneInterval(semi_distance)
		return output
	} else {
		return "Not enough notes were found to calculate a distance."
	}
}
