package main

import (
	"os"
)

func main() {
	//	s, sep := "", ""

	for index, arg := range os.Args[0:] {

		println(index, arg)

		//		s += sep + arg
		//sep = " "
	}

	// fmt.Println(s)
}
