package main

import (
	"fmt"
	"os"
)

type block interface {
	blockShow()
	getVal() string
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

func (a atom) getVal() string {
	return a.str
}

func (v variable) getVal() string {
	return v.str
}

func (s str) getVal() string {
	return s.str
}

func (o oParenth) getVal() string {
	return "("
}

func (c cParenth) getVal() string {
	return ")"
}

func (q quest) getVal() string {
	return "?"
}

func (p period) getVal() string {
	return "."
}

func (c col) getVal() string {
	return ":"
}

func (q eq) getVal() string {
	return "="
}

func (c compound) getVal() string {
	fmt.Println("whoops, getVal was called on compound type, something went wrong")
	os.Exit(3)
	return ""
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
		if b == nil {
			break
		}
		b.blockShow()
		//fmt.Printf("testing testing::: %T\n", c.body[i])
		//if fmt.Sprintf("%T", c.body[i]) == "*main.oParenth" {
		//	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		//}
		if fmt.Sprintf("%T", c.body[i]) != "*main.oParenth" && fmt.Sprintf("%T", c.body[i+1]) != "*main.cParenth" && fmt.Sprintf("%T", c.body[i]) != "*main.cParenth" && fmt.Sprintf("%T", c.body[i-1]) != "*main.cParenth" {
			fmt.Print(", ")
		}
	}
}
