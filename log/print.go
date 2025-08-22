// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package log

import (
	"os"

	"github.com/fatih/color"
)

// Prints a error message from a fatal error on the terminal using color red. Stops the process with code 1.
//
// Should be called whenever an error could cause unintended behaviour on a command.
func FatalError(err error) {
	red := color.New(color.FgRed).PrintlnFunc()
	red("Error:", err)
	os.Exit(1)
}

// Prints a error message from a non-fatal error on the terminal using color yellow.
//
// Should be called whenever an error should be acknowledged by the user.
func Warning(err error) {
	yellow := color.New(color.FgYellow).PrintlnFunc()
	yellow("Warning:", err)
}

// Prints the a string message on the terminal using color green. Stops the process with code 0.
// Should be called whenever a command reaches its intended goal or behavior, regardless of the variant flow or flags used.
func Success(succesMessage string) {
	yellow := color.New(color.FgGreen).PrintlnFunc()
	yellow(succesMessage)
	os.Exit(0)
}
