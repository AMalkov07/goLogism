package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

func parse(f *strings.Reader) {
	var tokens []block
	var s scanner.Scanner
	s.Init(f)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
		tmp := s.TokenText()
		switch tmp {
		case "(":
			tokens = append(tokens, (new(oParenth)))
		case ")":
			tokens = append(tokens, (new(cParenth)))
		case "?":
			tokens = append(tokens, (new(quest)))
		case ".":
			tokens = append(tokens, (new(period)))
		default:
			if unicode.IsUpper(rune(tmp[0])) {
				tokens = append(tokens, variable{str: tmp})
			} else if unicode.IsLower(rune(tmp[0])) {
				tokens = append(tokens, (atom{str: tmp}))
			} else if tmp[0] == '"' {
				tokens = append(tokens, str{str: tmp})
			}
		}
	}
	for _, val := range tokens {
		fmt.Printf("%T: ", val)
		val.blockShow()
		fmt.Println()
	}
}
