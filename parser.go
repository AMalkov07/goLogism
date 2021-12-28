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
			fmt.Println(tokens)
			tokens = append(tokens, (new(cParenth)))
			p := tokens[cStart[len(cStart)-1]]
			//fmt.Printf("is it a pointer???? well it is a: %T\n", p)
			ctmp := new(compound)
			ctmp.body = make([]block, 100)
			ctmp.head = p
			//fmt.Printf("is it a pointer???? well it is a: %T\n", p)
			copy(ctmp.body, tokens[cStart[len(cStart)-1]+1:len(tokens)])
			//copy_1 := copy(ctmp.body, tokens)
			//fmt.Println(copy_1)
			//fmt.Println(tokens)
			tokens[cStart[len(cStart)-1]] = ctmp
			tokens = tokens[:cStart[len(cStart)-1]+1]
			cStart = cStart[:len(cStart)-1]
			//tokens[len(cStart)-1] = compound{head: tokens[len(cStart)-1], body: make([]block, 0)}

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
