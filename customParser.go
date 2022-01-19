package main

import "fmt"

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
		token, ok := <-tokens
		if !ok {
			break
			//panic("no more tokens")
		}
		switch token.typ {
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
				com.outer = token
			} else if pCount == 1 {
				com.inner = append(com.inner, token)
			} else {
				tmp = &com
				for i := pCount; i > 1; i-- {
					tmp, ok = tmp.inner[len(tmp.inner)-1].(*compound)
					if !ok {
						panic("something went wrong with the compound")
					}
					tmp.inner = append(tmp.inner, token)
				}
			}
		case punct:
			if token.value == "?" {
				output = append(output, quer{com})
			} else {
				output = append(output, pred{com})
			}
			com = compound{}
		case comment:
			continue
		default:
			panic("something went wrong, couldn't identify token")
		}
	}
	return output
}
