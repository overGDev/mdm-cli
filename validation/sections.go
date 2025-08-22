// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package validation

import (
	"mdm/constants"
	"os"
)

// Checks if the sections folder exists on the current workDir.
func SectionsFolderExists() bool {
	_, err := os.Stat(constants.SECTIONS_FOLDER_NAME)
	return err == nil
}
