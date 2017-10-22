package gnirts

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// Reverse1 is a string reverser
func Reverse1(s string) string {
	// Credit: https://stackoverflow.com/a/1754209
	// Via: https://groups.google.com/forum/#!topic/golang-nuts/oPuBaYJ17t4
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

// Reverse2 is a string reverser
func Reverse2(s string) string {
	// Credit: https://stackoverflow.com/a/10030772
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Reverse3 is a string reverser
func Reverse3(s string) string {
	// Credit: https://play.golang.org/p/S2NIvT_DJ6
	// Via: A reply to https://stackoverflow.com/a/10030772
	c := []rune(s)
	n := len(c)
	for i := 0; i < n/2; i++ {
		c[i], c[n-1-i] = c[n-1-i], c[i]
	}
	return string(c)
}

// Reverse4 is a string reverser
func Reverse4(s string) string {
	// Credit: https://stackoverflow.com/a/4965535
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

// Reverse5 is a string reverser
func Reverse5(s string) string {
	// Credit: https://stackoverflow.com/a/4966500
	// compare this to Reverse8
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes)
}

// Reverse6 is a string reverser
func Reverse6(s string) string {
	// Credit: https://stackoverflow.com/a/20225618
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

// Reverse7 is a string reverser
func Reverse7(s string) string {
	// Credit: https://stackoverflow.com/a/26717253
	bound := make([]int, 0, len(s)+1)

	var iter norm.Iter
	iter.InitString(norm.NFD, s)
	bound = append(bound, 0)
	for !iter.Done() {
		iter.Next()
		bound = append(bound, iter.Pos())
	}
	bound = append(bound, len(s))
	out := make([]byte, 0, len(s))
	for i := len(bound) - 2; i >= 0; i-- {
		out = append(out, s[bound[i]:bound[i+1]]...)
	}
	return string(out)
}

// Reverse8 is a string reverser
func Reverse8(s string) string {
	// Credit: https://stackoverflow.com/a/1758098
	// Minor modifications to make it compile
	// compare this to Reverse5
	o := make([]rune, utf8.RuneCountInString(s))
	i := len(o)
	for _, c := range s {
		i--
		o[i] = c
	}
	return string(o)
}
