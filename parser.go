package main

import (
	"strings"
	"text/scanner"
	"unicode"
)

func compoundBodyCopy(c *compound, bs []block) {
	for _, b := range bs {
		c.body = append(c.body, b)
	}
}

func parse(f *strings.Reader) []stmt {
	var tokens []block
	var stmts []stmt
	var cStart []int
	var s scanner.Scanner
	s.Init(f)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tmp := s.TokenText()
		switch tmp {
		case "(":
			tokens = append(tokens, (new(oParenth)))
			cStart = append(cStart, len(tokens)-2)
		case ")":
			tokens = append(tokens, (new(cParenth)))
			p := tokens[cStart[len(cStart)-1]]
			ctmp := new(compound)
			ctmp.body = make([]block, 0)
			ctmp.head = p
			compoundBodyCopy(ctmp, tokens[cStart[len(cStart)-1]+1:])
			tokens[cStart[len(cStart)-1]] = ctmp
			tokens = tokens[:cStart[len(cStart)-1]+1]
			cStart = cStart[:len(cStart)-1]

		case "?":
			tokens = append(tokens, quest{})
			stmts = append(stmts, quer{body: tokens})
			tokens = nil
		case ".":
			tokens = append(tokens, period{})
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
	return stmts
}
