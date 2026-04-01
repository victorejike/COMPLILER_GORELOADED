package main

import (
	"fmt"
	"helper/helper"
	"os"
)

func main() {
	args := (os.Args)
	if len(args) != 3 {
		fmt.Println("Insufficient Commands")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fetch, _ := os.ReadFile(inputFile)

	SampleIn := string(fetch)

	SampleIn = helper.Upper(SampleIn)
	SampleIn = helper.ToLower(SampleIn)
	SampleIn = helper.FixAToAn(SampleIn)
	SampleIn = helper.ConvertNumbers(SampleIn)
	SampleIn = helper.ToCap(SampleIn)
	SampleIn = helper.BinTodecimal(SampleIn)
	SampleIn = helper.FixPunct(SampleIn)
	SampleIn = helper.FixSingleQuotes(SampleIn)
	// SampleIn = helper.FixQuotes(SampleIn)

	err := os.WriteFile(outputFile, []byte(SampleIn+"\n"), 0644)
	if err != nil {
		fmt.Println("error writing file")
		return
	}

	fmt.Println("Text tranformed succesfully")
	fmt.Println("Check your output manually or (cat output.txt)")

}
