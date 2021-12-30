package main

import (
	"fmt"
)

func blockCmp(b1 []block, b2 []block) bool {
	var type1 string
	var type2 string
	//fmt.Printf("size of first block is: %v\n", len(b1))
	for i, j := 0, 0; i < len(b1) && j < len(b2); {
		type1 = fmt.Sprintf("%T", b1[i])
		type2 = fmt.Sprintf("%T", b2[j])
		if type1 == "main.period" && type2 == "main.quest" || type1 == "main.quest" && type2 == "main.period" {
			return true
		}
		if type1 != type2 {
			return false
		}
		if type1 == "*main.compound" || type1 == "main.compound" {
			if b1[i].(*compound).head.getVal() != b2[j].(*compound).head.getVal() {
				return false
			}
			if !blockCmp(b1[i].(*compound).body, b2[j].(*compound).body) {
				return false
			} else {
				i++
				j++
				continue
			}
		}
		if b1[i].getVal() != b2[j].getVal() {
			return false
		}
		i++
		j++
	}
	return true
}

func evalQuer(s stmt, ss []stmt) {
	for _, elem := range ss {
		if blockCmp(s.(quer).body, elem.(pred).body) {
			fmt.Println("True")
			return
		}
	}
	fmt.Println("False")
}

func evaluate(s []stmt) {
	var preds []stmt
	for _, s := range s {
		if fmt.Sprintf("%T", s) == "main.pred" {
			preds = append(preds, s)
		} else if fmt.Sprintf("%T", s) == "main.quer" {
			evalQuer(s, preds)
		}
	}
	/*for _, elem := range preds {
		elem.stmtShow()
	}*/
}
