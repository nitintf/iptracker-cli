package print

import (
	"fmt"

	"github.com/fatih/color"
)

// ErrorPrint prints an error message in red color.
func Error(message string) {
	red := color.New(color.FgHiRed).Add(color.Bold)
	red.Println(message)
}

// HeadingPrint prints a heading in cyan color with underline.
func Heading(message string) {
	cyan := color.New(color.FgCyan)
	cyan.Add(color.Underline)
	cyan.Println(message)
}

// Print prints text in the default color.
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Info(message string) {
	c := color.New(color.FgGreen)
	c.Println(message)
}
