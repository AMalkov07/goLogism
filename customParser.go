package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type tokenType int // our tokens are basically just ghetto enums
const (
	oParen tokenType = iota // iota is a magic word that makes our values equal to number. (first value in struct is 0, next is 1, etc)
	cParen
	punct
	comma
	atom
	variable
	str
)

type token struct {
	typ   tokenType
	value string
}

type lexer struct {
	input  string
	start  int
	pos    int
	width  int        // might not need, this is the width of the current character we are looking at (number of bytes of current character)
	tokens chan token // this being a chan allows us to have the lexer running and creating toeksn on 1 thread, and a parser actually parsing those tokens at the same time on another thread
}

//func parse(tokens chan token) Node {
func parse(tokens chan token) {
	for {
		token, ok := <-tokens
		if !ok {
			panic("no more tokens")
		}
		fmt.Print(token.value, ",")
	}
}

const eof rune = -1 // this eof will tell us when we are done w/ the whole file

type stateFunc func(*lexer) stateFunc

//func BeginLexing(s string) Node {
func BeginLexing(s string) {
	l := &lexer{
		input:  s,
		tokens: make(chan token, 100)} // we want the channel to be buffered so that the lexer can keep adding tokens, even if the parser hasn't pulled them out to parse them yet
	go l.run() // go routine
	parse(l.tokens)
}

func (l *lexer) run() {
	for state := determineToken; state != nil; { // this loop starts w/ state being set to our determineToken function, and then continues running till state is nil
		state = state(l) // through each iteration of the looop state is equal to the state it gets back from the current state, when we pass in the current lexer
	}
	close(l.tokens) // close the channel
}

func determineToken(l *lexer) stateFunc {
	for {
		switch r := l.next(); {
		case isWhiteSpace(r):
			l.ignore()
		case r == '(':
			l.emit(oParen)
		case r == ')':
			l.emit(cParen)
		case r == ',':
			l.emit(comma)
		case isUpper(r): // means that we are at the start of a var
			return lexVar
		case isLower(r): // means that we are at the start of a Atom
			return lexAtom
		case r == '"': // means that we are at the start of a str
			return lexStr
		case isPunct(r):
			return lexPunct
		case r == eof:
			return nil
		default:
			panic("unknown input was discovered")
			//fmt.Println("unknown input was discovered")
			//os.Exit(3)
		}
	}
}

func lexPunct(l *lexer) stateFunc {
	l.accept(".?")
	l.emit(punct)
	return determineToken
}

func lexVar(l *lexer) stateFunc {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	l.accept(upper)
	l.acceptRun(lower)
	l.emit(variable)
	return determineToken
}

func lexAtom(l *lexer) stateFunc {
	lower := "abcdefghijklmnopqrstuvwxyz"
	l.accept(lower)
	l.acceptRun(lower)
	l.emit(atom)
	return determineToken
}

func lexStr(l *lexer) stateFunc {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	l.accept("\"")
	l.acceptRun(letters)
	l.accept("\"")
	l.emit(str)
	return determineToken
}

func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isPunct(r rune) bool {
	return r == '.' || r == '?'
}

func (l *lexer) emit(t tokenType) {
	l.tokens <- token{t, l.input[l.start:l.pos]} // we are writing a new token to the tokens channell
	l.start = l.pos
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) { // checks if we are done w/ the file
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:]) // we are taking a slice of our string that goes from our lexer position to the end, and the DecodeRuneInString gets the next rune in that input and returns the width of that rune
	l.pos += l.width                                      // our position will be updated w/ the width of the rune we just got in above step
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

/* usefull but not needed functions:
func (l *lexer) peek() (r rune) {
	r; _ = utf8.DecodeRuneInString(l.input[l.pos:])
	return r
}
*/
