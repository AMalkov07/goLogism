package main

import "fmt"

func main() {
	var iTest []stmt
	aTest := atom{str: "this is atom str"}
	vTest := variable{str: "this is variable str"}
	cTest2 := compound{head: "head2",
		body: vTest}
	cTest := compound{head: "head1",
		body: cTest2}
	qTest := quer{body: aTest}
	pTest := pred{body: cTest}
	iTest = append(iTest, qTest)
	iTest = append(iTest, pTest)
	for _, e := range iTest {
		e.stmtShow()
		fmt.Println()
	}
}
