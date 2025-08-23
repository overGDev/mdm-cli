// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package application

import (
	"mdm/internal/templates"
	"os"
	"path"
)

const (
	SCRIPT_FOLDER_NAME = "scripts"
	SCRIPT_FILE_NAME   = "combine_sections.py"
)

// Creates a schema file with default sample content.
// The content is intended to showcase the customization the tool provides.
func GenerateScript() error {
	err := os.Mkdir(SCRIPT_FOLDER_NAME, 0644)
	if err != nil {
		return err
	}

	file := path.Join(SCRIPT_FOLDER_NAME, SCRIPT_FILE_NAME)
	return os.WriteFile(file, []byte(templates.Script), 0644)
}
