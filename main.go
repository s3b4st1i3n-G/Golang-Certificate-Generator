package main

import (
	"fmt"
	"gencert/cert"
	"gencert/pdf"
	"os"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Johnson", "2023-11-27")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during pdf generation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)
}
