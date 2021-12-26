package main

import "fmt"

type block interface {
	blockShow()
}

type atom struct {
	str string
}

type variable struct {
	str string
}

type str struct {
	str string
}

type oParenth struct {
}

type cParenth struct {
}

type quest struct {
}

type period struct {
}

type col struct {
}

type eq struct {
}

type compound struct {
	head block
	body []block
}

func (a atom) blockShow() {
	fmt.Print(a.str)
}

func (v variable) blockShow() {
	fmt.Print(v.str)
}

func (s str) blockShow() {
	fmt.Print(s.str)
}

func (o oParenth) blockShow() {
	fmt.Printf("%c", '(')
}

func (c cParenth) blockShow() {
	fmt.Printf("%c", ')')
}

func (q quest) blockShow() {
	fmt.Printf("%c", '?')
}

func (p period) blockShow() {
	fmt.Printf("%c", '.')
}

func (c col) blockShow() {
	fmt.Printf("%c", ':')
}

func (q eq) blockShow() {
	fmt.Printf("%c", '=')
}

func (c compound) blockShow() {
	c.head.blockShow()
	for i, b := range c.body {
		b.blockShow()
		if fmt.Sprintf("%T", c.body[i+1]) != "main.cParenth" {
			fmt.Print(", ")
		}
	}
}
