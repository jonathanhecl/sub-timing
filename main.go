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

	mode := ""
	firstLine := ""
	lastLine := ""
	shift := ""

	// Parse command line arguments
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "-s=") {
			source = strings.TrimPrefix(arg, "-s=")
		} else if strings.HasPrefix(arg, "-t=") {
			target = strings.TrimPrefix(arg, "-t=")
		} else if strings.HasPrefix(arg, "-m=") {
			mode = strings.TrimPrefix(arg, "-m=")
		} else if strings.HasPrefix(arg, "-f=") {
			firstLine = strings.TrimPrefix(arg, "-f=")
		} else if strings.HasPrefix(arg, "-l=") {
			lastLine = strings.TrimPrefix(arg, "-l=")
		} else if strings.HasPrefix(arg, "-d=") {
			shift = strings.TrimPrefix(arg, "-d=")
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

	if mode == "" {
		fmt.Println("Error: Mode not specified (move, shift, adjust)")
		return
	} else if mode != "move" && mode != "shift" && mode != "adjust" {
		fmt.Println("Error: Invalid mode (move, shift, adjust)")
		return
	}

	if mode == "move" || mode == "adjust" {
		if firstLine == "" {
			fmt.Println("Error: First line not specified (-f)")
			return
		}
	}

	if mode == "adjust" {
		if lastLine == "" {
			fmt.Println("Error: Last line not specified (-l)")
			return
		}
	}

	if mode == "shift" {
		if shift == "" {
			fmt.Println("Error: Shift duration not specified (-d)")
			return
		}
	}

	fmt.Println("Source:", source)
	fmt.Println("Target:", target)
	fmt.Println("Mode:", mode)
	if mode == "move" || mode == "adjust" {
		fmt.Println("First Line:", firstLine)
	}
	if mode == "adjust" {
		fmt.Println("Last Line:", lastLine)
	}
	if mode == "shift" {
		fmt.Println("Shift:", shift)
	}

	// Load source subtitle file
	sub := subtitles.Subtitle{}
	err := sub.LoadFile(source)
	if err != nil {
		fmt.Printf("Error loading subtitle file: %v\n", err)
		return
	}

	fmt.Printf("Total lines: %d\n\n", len(sub.Lines))
}
