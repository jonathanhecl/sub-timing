package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jonathanhecl/subtitle-processor/subtitles"
)

var (
	version = "1.0.0"
	source  = ""
	target  = ""
)

func main() {
	fmt.Println("Sub-Timing v" + version)

	// Parse command line arguments
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "-s=") {
			source = strings.TrimPrefix(arg, "-s=")
		} else if strings.HasPrefix(arg, "-t=") {
			target = strings.TrimPrefix(arg, "-t=")
		}
	}

	if source == "" {
		fmt.Println("Error: Source file not specified")
		return
	}

	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Printf("Source file %s does not exist\n", source)
		return
	}

	if target == "" {
		// Auto-generate target filename
		ext := path.Ext(source)
		target = strings.TrimSuffix(source, ext) + "_modified" + ext
	}

	fmt.Println("Source:", source)
	fmt.Println("Target:", target)

	// Load source subtitle file
	sub := subtitles.Subtitle{}
	err := sub.LoadFile(source)
	if err != nil {
		fmt.Printf("Error loading subtitle file: %v\n", err)
		return
	}

	fmt.Printf("Total lines: %d\n\n", len(sub.Lines))
}
