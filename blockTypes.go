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

type compound struct {
	head string
	body block
}

func (a atom) blockShow() {
	fmt.Print(a.str)
}

func (v variable) blockShow() {
	fmt.Print(v.str)
}

func (c compound) blockShow() {
	fmt.Printf("%v(", c.head)
	c.body.blockShow()
	fmt.Print(")")
}
