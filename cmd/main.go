package main

import (
	"archiver/internal/archiver"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileName := flag.String("f", "", "file with extension")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help || *fileName == "" {
		fmt.Println("use -f <file> to choose file")
		return
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("error while opening file: ", err.Error())
		return
	}
	defer file.Close()

	fName := strings.ReplaceAll(filepath.Base(file.Name()), filepath.Base(filepath.Ext(file.Name())), ".hf")
	newFile, err := os.Create(fName)
	if err != nil {
		fmt.Printf("error while creating %s: %s\n", fName, err.Error())
	}
	defer newFile.Close()

	err = archiver.Huffman(file, newFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
