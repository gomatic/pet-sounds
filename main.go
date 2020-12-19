package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	defaultFileName = "pets.hcl"
)

func main() {
	if err := inner(); err != nil {
		fmt.Printf("pet-sounds error: %s\n", err.Error())
		os.Exit(1)
	}
}

func inner() error {

	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{defaultFileName}
	}

	// There is a random function for the HCL configuration.
	rand.Seed(time.Now().Unix())

	for _, inputFile := range args {
		log.Print(inputFile)
		pets, err := ReadConfig(inputFile)
		if err != nil {
			return err
		}

		for _, p := range pets {
			p.Say()
			p.Act()
		}
	}

	return nil
}
