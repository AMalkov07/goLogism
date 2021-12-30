package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("plz add input")
		os.Exit(3)
	}
	ff, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	ff = ff[2:]
	for i := 0; i < len(ff); {
		if ff[i] == 0 {
			ff = append(ff[:i], ff[i+1:]...)
		} else {
			i++
		}
	}
	//fmt.Println((ff))
	//f := os.Args[1]
	//f := strings.NewReader("//testing\n\"parent\"(Alexei, Olga).\nparent(Alexei, Andrey)?")
	evaluate(parse(strings.NewReader(string(ff))))
}
