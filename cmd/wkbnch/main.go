package main

import (
	"flag"
	_ "fmt"
	"log"
)

type Config struct {
	In   *string
	Out  *string
	Lang *string
}

func main() {
	// configure
	config := Config{
		In:   flag.String("in", "", "input file or directory"),
		Out:  flag.String("out", "", "output directory"),
		Lang: flag.String("lang", "c", "language to generate"),
	}

	flag.Parse()

	if *config.In == "" {
		log.Fatalln("input file or directory must be defined with the -in flag")
	}

	if *config.Out == "" {
		log.Fatalln("output dirctory must be defined with the -out flag")
	}

	// find files
	files, err := filesFromFlag(*config.In)
	if err != nil {
		log.Fatalf("failed getting files from input given in flag -in: %v", err)
	}

	objects, err := objectsFromFilenames(files)
	if err != nil {
		log.Fatalf("failed parsing in files: %v", err)
	}

	log.Printf("parsed in %d objects from %d files", len(objects), len(files))

	// generate code

	// output
}
