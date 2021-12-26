package main

import "fmt"

type stmt interface {
	stmtShow()
}

type pred struct {
	body []block
}
type quer struct {
	body []block
}

func (q quer) stmtShow() {
	for _, b := range q.body {
		b.blockShow()
	}
	fmt.Println()
}

func (p pred) stmtShow() {
	for _, b := range p.body {
		b.blockShow()
	}
	fmt.Println()
}
