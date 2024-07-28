package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AlekseiAQ/command-line-utility-with-flags/pkg/tools"
)

var (
	isRecursive bool
	isHuman     bool
)

func init() {
	flag.BoolVar(&isRecursive, "recursive", false, "Calculate sizes recursively")
	flag.BoolVar(&isHuman, "human", false, "Display sizes in human-readable format")
}

func main() {
	flag.Parse()
	directories := flag.Args()

	if len(directories) == 0 {
		fmt.Println("Please provide at least one directory")
		os.Exit(1)
	}

	var total int64
	for _, d := range directories {
		size := tools.CalculateDirSize(d, isRecursive, isHuman)
		fmt.Printf("%s: %s\n", d, tools.FormatSize(size, isHuman))
		total += size
	}

	fmt.Printf("Total: %s\n", tools.FormatSize(total, isHuman))
}
