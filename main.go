package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Define variables for the command line
	var (
		file    = flag.String("file", "", "file name")
		search  = flag.String("search", "", "'search string'")
		replace = flag.String("replace", "", "'replace string'")
	)

	flag.Parse()

	// Check input file
	if len(*file) == 0 {
		fmt.Println("Will not specify an input file")
		os.Exit(1)
	}

	// Check input search string
	if len(*search) == 0 {
		fmt.Println("Will not specify an search string")
		os.Exit(1)
	}

	// Check input replace string
	if len(*replace) == 0 {
		fmt.Println("Will not specify an replace string")
		os.Exit(1)
	}

	//Reading input file
	input, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Finding a replacement string
	if !bytes.Contains(input, []byte(*search)) {
		fmt.Println("Search string not found")
		os.Exit(1)
	}

	//Replacing string
	output := bytes.Replace(input, []byte(*search), bytes.TrimSpace([]byte(*replace)), -1)
	fmt.Println("Replacing string is complete")

	//Get file permission
	fileInfo, _ := os.Stat(*file)
	mode := fileInfo.Mode()

	//Writing file
	if err = ioutil.WriteFile(*file, output, mode); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
