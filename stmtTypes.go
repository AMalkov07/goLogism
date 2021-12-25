package main

type stmt interface {
	stmtShow()
}

type pred struct {
	body block
}
type quer struct {
	body block
}

func (q quer) stmtShow() {
	q.body.blockShow()
}

func (p pred) stmtShow() {
	p.body.blockShow()
}
