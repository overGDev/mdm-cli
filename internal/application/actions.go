// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package application

import (
	"mdm/internal/templates"
	"os"
	"path"
)

const (
	GITHUB_FOLDER_NAME         = ".github"
	GITHUB_ACTIONS_FOLDER_NAME = "workflows"
	GITHUB_ACTION_FILE_NAME    = "update_readme.yaml"
)

// Creates the GitHub Action file necesary to combine all of the document sections
// on a single markdown file. The .yaml file is created inside the necesary folder path.
func GenerateGitHubAction() error {
	folder := path.Join(GITHUB_FOLDER_NAME, GITHUB_ACTIONS_FOLDER_NAME)

	err := os.MkdirAll(folder, 0644)
	if err != nil {
		return err
	}

	file := path.Join(folder, GITHUB_ACTION_FILE_NAME)
	return os.WriteFile(file, []byte(templates.GitHubAction), 0644)
}
