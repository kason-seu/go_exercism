package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	cs := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}
	for _, r := range log {
		if v, ok := cs[r]; ok {
			return v
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	// rs := ""

	// for _, value := range log {
	// 	if value == oldRune {
	// 		rs += string(newRune)
	// 	} else {
	// 		rs += string(value)
	// 	}

	// }
	// return rs

    // vmethod 2
    // var rs []rune
    // for _, value := range log{
    //     if value == oldRune {
    //         rs = append(rs, newRune)
    //     } else {
    //         rs = append(rs, value)
    //     }
    // }
    // return string(rs)

    // method 3
    var rs []rune = []rune(log)
    for i, value := range rs {
        if value == oldRune {
            rs[i] = newRune
        }
    }
    return string(rs)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	return utf8.RuneCountInString(log) <= limit

}
