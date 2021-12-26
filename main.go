package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("plz add input")
	}
	/*f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}*/
	f := strings.NewReader("//testing\n\"parent\"(Alexei, Olga).")
	parse(f)
}
