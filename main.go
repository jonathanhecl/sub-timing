package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

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
	firstLine := time.Duration(0)
	lastLine := time.Duration(0)
	shift := time.Duration(0)
	negativeShift := false

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
			firstLine, _ = parseDuration(strings.TrimPrefix(arg, "-f="))
		} else if strings.HasPrefix(arg, "-l=") {
			lastLine, _ = parseDuration(strings.TrimPrefix(arg, "-l="))
		} else if strings.HasPrefix(arg, "-d=") {
			if strings.HasPrefix(strings.TrimPrefix(arg, "-d="), "-") {
				negativeShift = true
			}
			shift, _ = parseDuration(strings.TrimPrefix(arg, "-d="))
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
		if firstLine == time.Duration(0) {
			fmt.Println("Error: First line not specified (-f)")
			return
		}
	}

	if mode == "adjust" {
		if lastLine == time.Duration(0) {
			fmt.Println("Error: Last line not specified (-l)")
			return
		}
	}

	if mode == "shift" {
		if shift == time.Duration(0) {
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

	switch mode {
	case "move":
		sub = subMove(sub, firstLine)
	case "shift":
		sub = subShift(sub, shift, negativeShift)
	case "adjust":
		// sub = subAdjust(sub, firstLine, lastLine)
	}

	// Save modified subtitle file
	err = sub.SaveFile(target)
	if err != nil {
		fmt.Printf("Error saving subtitle file: %v\n", err)
		return
	}

	fmt.Println("\nDone.")

}

// Parse duration string into time.Duration
// Example: "0:00:00.000" (h:mm:ss.mmm) -> 0
func parseDuration(duration string) (time.Duration, error) {
	duration = strings.TrimSpace(duration)

	// Handle empty string
	if duration == "" {
		return 0, nil
	}

	// Check if the format is already compatible with time.ParseDuration
	if strings.Contains(duration, "h") || strings.Contains(duration, "m") || strings.Contains(duration, "s") {
		return time.ParseDuration(duration)
	}

	// Split the time components
	parts := strings.Split(duration, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid time format: %s, expected h:mm:ss.mmm", duration)
	}

	// Parse hours
	hours, err := time.ParseDuration(parts[0] + "h")
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %s", parts[0])
	}

	// Parse minutes
	minutes, err := time.ParseDuration(parts[1] + "m")
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %s", parts[1])
	}

	// Parse seconds (may include milliseconds)
	seconds, err := time.ParseDuration(parts[2] + "s")
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %s", parts[2])
	}

	// Combine all parts
	return hours + minutes + seconds, nil
}

func subMove(sub subtitles.Subtitle, firstLine time.Duration) subtitles.Subtitle {
	ret := sub

	diff := ret.Lines[0].Start - firstLine

	ret.Lines[0].Start = firstLine
	ret.Lines[0].End = ret.Lines[0].End - diff

	for i := 1; i < len(ret.Lines); i++ {
		ret.Lines[i].Start = ret.Lines[i].Start - diff
		ret.Lines[i].End = ret.Lines[i].End - diff
	}
	return ret
}

func subShift(sub subtitles.Subtitle, shift time.Duration, negativeShift bool) subtitles.Subtitle {
	ret := sub

	for i := 0; i < len(ret.Lines); i++ {
		if negativeShift {
			ret.Lines[i].Start = ret.Lines[i].Start - shift
			ret.Lines[i].End = ret.Lines[i].End - shift
		} else {
			ret.Lines[i].Start = ret.Lines[i].Start + shift
			ret.Lines[i].End = ret.Lines[i].End + shift
		}
	}
	return ret
}
