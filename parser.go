package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

func parse(f *strings.Reader) {
	var tokens []block
	var stmts []stmt
	var cStart []int
	var s scanner.Scanner
	s.Init(f)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
		tmp := s.TokenText()
		switch tmp {
		case "(":
			tokens = append(tokens, (new(oParenth)))
			cStart = append(cStart, len(tokens)-2)
		case ")":
			tokens = append(tokens, (new(cParenth)))
			tokens[len(cStart)-1] = compound{head: tokens[len(cStart)-1], body: make([]block, 0)}
		case "?":
			tokens = append(tokens, (new(quest)))
			stmts = append(stmts, quer{body: tokens})
			tokens = nil
		case ".":
			tokens = append(tokens, (new(period)))
			stmts = append(stmts, pred{body: tokens})
			tokens = nil
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
	for _, val := range stmts {
		fmt.Printf("%T: ", val)
		val.stmtShow()
	}
}
