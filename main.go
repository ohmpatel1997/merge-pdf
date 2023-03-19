package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	// Define command-line flags
	outputFilename := flag.String("output", "merged.pdf", "the name of the output PDF file")
	flag.Parse()

	// Load the input PDF files
	inputFiles := flag.Args()

	if len(inputFiles) < 2 {
		fmt.Println("Error: At least two input PDF files are required")
		os.Exit(1)
	}

	// Define the PDF processing configuration
	config := model.NewDefaultConfiguration()

	// Merge the PDF files
	err := api.MergeCreateFile(inputFiles, *outputFilename, config)
	if err != nil {
		fmt.Println("Error merging PDF files:", err)
		os.Exit(1)
	}

	fmt.Printf("PDF files merged successfully. Output file: %s\n", *outputFilename)
}
