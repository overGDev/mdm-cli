// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package application

import (
	"mdm/internal/templates"
	"os"
)

const SCHEMA_FILE_NAME = "schema.yaml"

// Checks if the schema file exists on the current workDir.
func SchemaFileExists() bool {
	_, err := os.Stat(SCHEMA_FILE_NAME)
	return err == nil
}

// Creates a schema file with default sample content.
// The content is intended to showcase the customization the tool provides.
func GenerateSampleSchema() error {
	return os.WriteFile(SCHEMA_FILE_NAME, []byte(templates.SampleSchema), 0644)
}
