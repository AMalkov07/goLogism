package main

import (
	"fmt"
	"strconv"
)

type stmt interface {
	showInterface()
}

type pred struct {
	c compound
}

type quer struct {
	c compound
}

func (p pred) showInterface() {
	p.c.show()
	fmt.Println(".")
}

func (q quer) showInterface() {
	q.c.show()
	fmt.Println("?")
}

func parse(tokens chan token) []stmt {
	output := []stmt{}
	com := compound{}
	var tmp *compound
	pCount := 0
	for {
		tk, ok := <-tokens
		//fmt.Printf("token type is: %v\n", token.typ)
		if !ok {
			break
			//panic("no more tokens")
		}
		switch tk.typ {
		case oParen:
			pCount++
			if pCount > 1 {
				com.inner[len(com.inner)-1] = &compound{outer: com.inner[len(com.inner)-1]}
			}
		case cParen:
			pCount--
		//case comma:
		case atom, variable, str, comma:
			if pCount == 0 {
				com.outer = tk
				/*} else if pCount == 1 {
				com.inner = append(com.inner, tk)*/
			} else if pCount == 1 {
				com.inner = append(com.inner, tk)
			} else {
				tmp = &com
				for i := pCount; i > 1; i-- {
					tmp, ok = tmp.inner[len(tmp.inner)-1].(*compound)
					if !ok {
						panic("something went wrong with the compound type")
					}
					tmp.inner = append(tmp.inner, tk)
				}
			}
		case shortcut:
			if pCount == 0 {
				panic("misplaced shortcut")
			} else if pCount == 1 {
				if tk.value == "0" {
					com.inner = append(com.inner, token{typ: atom, value: "zero"})
				}
				com.inner = append(com.inner, tk.stoc())
			} else {
				tmp = &com
				for i := pCount; i > 1; i-- {
					tmp, ok = tmp.inner[len(tmp.inner)-1].(*compound)
					if !ok {
						panic("something went wrong with the compound type")
					}
					tmp.inner = append(tmp.inner, tk.stoc())
				}
			}
		case punct:
			if tk.value == "?" {
				output = append(output, quer{com})
			} else {
				output = append(output, pred{com})
			}
			com = compound{}
		case comment:
			continue
		default:
			panic("parser couldn't identify token")
		}
	}
	return output
}

func (t token) stoc() *compound {
	counter, err := strconv.Atoi(t.value)
	if err != nil {
		panic("error happened in the stoc() function")
	}
	output := &compound{outer: token{typ: atom, value: "succ"}}
	trav := output
	for ; counter > 1; counter-- {
		trav.inner = append(trav.inner, &compound{outer: token{typ: atom, value: "succ"}})
		var ok bool
		trav, ok = trav.inner[0].(*compound)
		if !ok {
			panic("compound problem in stoc function")
		}
	}
	trav.inner = append(trav.inner, token{typ: atom, value: "zero"})
	return output
}
