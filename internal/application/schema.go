// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package application

import (
	"os"
)

const SCHEMA_FILE_NAME = "schema.yaml"
const SAMPLE_SCHEMA_CONTENT = `- title: Summary
- title: Installation and usage
  alias: installation
- title: System Architecture
  alias: Architecture
- title: Components
  children:
    - title: Backend
    - title: Frontend
    - title: Database`

// Checks if the schema file exists on the current workDir.
func SchemaFileExists() bool {
	_, err := os.Stat(SCHEMA_FILE_NAME)
	return err == nil
}

// Creates a schema file with default sample content.
// The content is intended to showcase the customization the tool provides.
func GenerateSampleSchema() error {
	return os.WriteFile(SCHEMA_FILE_NAME, []byte(SAMPLE_SCHEMA_CONTENT), 0644)
}
