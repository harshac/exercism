//Package raindrops to convert numbers to raindrop sounds
package raindrops

import "strconv"

//Convert converts a number to it's respective raindrop sound
func Convert(number int) string {
	var sound string
	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(number)
	}
	return sound
}
