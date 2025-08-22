// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package errors

import "fmt"

func ExistingFolderError(folderName string) error {
	return fmt.Errorf("'%s' folder already exists", folderName)
}

func ExistingFileError(fileName string) error {
	return fmt.Errorf("'%s' file already exists", fileName)
}

func MissingFolderError(folderName string) error {
	return fmt.Errorf("'%s' folder not found", folderName)
}

func MissingFileError(fileName string) error {
	return fmt.Errorf("'%s' file not found", fileName)
}
