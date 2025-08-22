// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package errors

import (
	"fmt"
	"mdm/constants"
)

func SchemaFileNotFound() error {
	return fmt.Errorf("'%s' file not found. Try using '--sample' to generate a sample schema", constants.SCHEMA_FILE_NAME)
}
