package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Includer stores context of includes
type Includer struct {
	includedirs []string
	dirContens  map[string]map[string]bool
	fileDeps    map[string][]string
}

// NewIncluder creates includer
func NewIncluder(pathes []string) *Includer {
	includer := Includer{pathes, make(map[string]map[string]bool), make(map[string][]string)}

	// read directory contencs of all include dirs once
	for _, i := range pathes {
		var err error
		f, err := os.Open(i)
		if err != nil {
			panic("could not open dir " + i)
		}
		defer f.Close()

		includer.dirContens[i] = make(map[string]bool)

		filelist, _ := f.Readdirnames(-1)
		for _, de := range filelist {
			includer.dirContens[i][de] = true
		}

		if err != nil {
			panic("error in reading dir " + i)
		}
	}

	return &includer
}

// ProcessFile ...
func (i *Includer) ProcessFile(file string) []string {
	dependencies := make([]string, 0, 1024)

	f, err := os.Open(file)
	if err != nil {
		panic("could not open file " + file)
	}
	defer f.Close()

	fmt.Println("reading ", file)

	regex, _ := regexp.Compile(`^\s*#\s*include\s*([<|"])(.*)[>|"]`)

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		match := regex.FindSubmatch(line)
		if match != nil {
			// delim := string(match[1])
			includename := string(match[2])
			fmt.Println(includename)

			for _, ii := range i.includedirs {
				_, ok := i.dirContens[ii][includename]
				if ok {
					dependencies = append(dependencies, ii+"/"+includename)
				}
			}
		}
	}

	return dependencies
}
