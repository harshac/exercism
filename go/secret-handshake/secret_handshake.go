package secret

var codes = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

const reverseCommand = 16

//Handshake converts decimal number to the appropriate sequence of events for a secret handshake
func Handshake(num uint) []string {
	var bases = []uint{1, 2, 4, 8}
	var events []string
	for _, k := range bases {
		if num&k == k && num&k != reverseCommand {
			events = append(events, codes[k])
		}
	}
	if num&reverseCommand == reverseCommand {
		return reverse(events)
	}
	return events
}

func reverse(codes []string) []string {
	var reversed []string
	for _, code := range codes {
		reversed = append([]string{code}, reversed...)
	}
	return reversed
}
