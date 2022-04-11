package main

import (
	"fmt"
	brainfuck "github.com/denizgursoy/brainfuck/pkg/domain"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 4 {
		log.Fatalln("you must specify 3 paths for files: command file, input file, output file respectively")
	}

	commandFile, commandFileError := os.Open(os.Args[1])
	if commandFileError != nil {
		log.Fatalln("could not read from command file")
	}
	inputFile, inputFileError := os.Open(os.Args[2])
	if inputFileError != nil {
		log.Fatalln("could not read from input file")
	}
	outputFile, outputFileError := os.OpenFile(os.Args[3], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if outputFileError != nil {
		log.Fatalln("could not read from output file")
	}

	defer commandFile.Close()
	defer inputFile.Close()
	defer outputFile.Close()

	options := brainfuck.IoOptions{
		CommandReader: commandFile,
		InputReader:   inputFile,
		OutputWriter:  outputFile,
	}

	fuck, err := brainfuck.NewBrainFuck(&options)
	if err != nil {
		log.Fatalln(err)
	}

	err = fuck.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
