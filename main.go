package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"path/filepath"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	home := os.Getenv("HOME")

	// Define command-line flags
	outputFilename := flag.String("output", "merged.pdf", "the name of the output PDF file")
	flag.Parse()

	*outputFilename = filepath.FromSlash( filepath.Clean(fmt.Sprintf("%s/%s", home, *outputFilename)))
	
	// Load the input PDF files
	inputFiles := flag.Args()

	if len(inputFiles) < 2 {
		fmt.Println("Error: At least two input PDF files are required")
		os.Exit(1)
	}

	// Define the PDF processing configuration
	config := model.NewDefaultConfiguration()


	for i := range inputFiles {
		inputFiles[i] = filepath.FromSlash( filepath.Clean(fmt.Sprintf("%s/%s", home, inputFiles[i])))
		fmt.Println("path--", inputFiles[i])
	}

	// Merge the PDF files
	err := api.MergeCreateFile(inputFiles, *outputFilename, config)
	if err != nil {
		fmt.Println("Error merging PDF files:", err)
		os.Exit(1)
	}

	fmt.Printf("PDF files merged successfully. Output file: %s\n", *outputFilename)
}
