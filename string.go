package cbutil

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringPrefix(prefix string) string {
	return strings.Join([]string{prefix, RandString(10)}, "")
}

func StringReplacer(format string, args ...string) string {
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}

func StringFirstLetterLower(s string) string {
	if len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r != utf8.RuneError || size > 1 {
			lo := unicode.ToLower(r)
			if lo != r {
				s = string(lo) + s[size:]
			}
		}
	}
	return s
}
