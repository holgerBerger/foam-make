package main

import (
	"fmt"
	"os"
)

func main() {

	includedirs := []string{"/usr/include"}

	includer := NewIncluder(includedirs)

	for _, file := range os.Args[1:] {
		fmt.Println(includer.ProcessFile(file))
	}

}
