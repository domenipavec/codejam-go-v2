package main

import "github.com/matematik7/codejam-go-v2/io"

func main() {
	io.TestCases(testCase)
}

func testCase(input *io.Input, output *io.Output) {
	output.Print(input.Int())
}
