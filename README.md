# PDF Merge Tool

This is a command-line tool for merging multiple PDF files into a single PDF file using Go programming language and the unidoc/pdf library.


## Installation

There are two way to install this tool:

### Homebrew

You can install this tool via homebrew using command:

     `brew install merge-pdf`

### Build from soures

1. Install Go: You can download and install it from the official website.

2. Install the unidoc/pdf library: You can install it using the following command:

     
    `go get github.com/unidoc/unidoc/pdf`


3. Clone this repository:

4. Navigate to the directory containing the cloned repository.

## Usage

`merge-pdf -output output.pdf input1.pdf input2.pdf input3.pdf`

Replace `input1.pdf`, `input2.pdf`, and `input3.pdf` with the names of the PDF files you want to merge, and replace output.pdf with the name you want to give to the merged output file.

## Contributing

Contributions are welcome! If you have any suggestions or bug reports, please create an issue or a pull request.


