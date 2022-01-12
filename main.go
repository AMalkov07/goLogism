package main

import (
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("plz add input")
	}
	args := os.Args
	f, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	fStr := string(f)
	BeginLexing(fStr)

	/*
		f = f[2:]
		for i := 0; i < len(ff); {
			if f[i] == 0 {
				ff = append(f[:i], f[i+1:]...)
			} else {
				i++
			}
		}*/
	//fmt.Println((ff))
	//f := os.Args[1]
	//f := strings.NewReader("//testing\n\"parent\"(Alexei, Olga).\nparent(Alexei, Andrey)?")
	//evaluate(parse(strings.NewReader(string(ff))))
}
