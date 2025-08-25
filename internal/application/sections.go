// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package application

import (
	"fmt"
	"io/fs"
	"mdm/internal/model"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
	"gopkg.in/yaml.v3"
)

const SECTIONS_FOLDER_NAME = "sections"

var consideredSections = []string{}

// Checks if the sections folder exists on the current workDir.
func SectionsFolderExists() bool {
	_, err := os.Stat(SECTIONS_FOLDER_NAME)
	return err == nil
}

// Reads from the 'schema.yaml' file, converting the data into file and folder representing each document section.
func LoadSchema() ([]model.Section, error) {
	data, err := os.ReadFile(SCHEMA_FILE_NAME)
	if err != nil {
		return nil, err
	}

	var sections []model.Section
	if err := yaml.Unmarshal(data, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}

// GenerateDocumentSections creates the folder and Markdown file structure
// from the given list of sections. Each section is represented either as a
// directory (if it has children) or as a Markdown file (if it is a leaf).
//
// The recursive expansion is delegated to expandSectionsTree.
func GenerateDocumentSections(sections []model.Section) error {
	baseDir := SECTIONS_FOLDER_NAME
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return err
	}

	for _, section := range sections {
		if err := expandSectionsTree(section, baseDir, 1); err != nil {
			return err
		}
	}

	return nil
}

// expandSectionsTree recursively builds the section hierarchy in the file
// system. Directories are created for parent sections, while leaf sections
// generate Markdown files with headers matching their depth.
func expandSectionsTree(section model.Section, parentPath string, headerLevel int) error {
	name := section.Title
	if section.Alias != "" {
		name = section.Alias
	}
	name = sanitizeString(name)

	// Recursive call in case it is a parent section with children
	if len(section.Children) > 0 {
		dirPath := filepath.Join(parentPath, name)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}

		for _, child := range section.Children {
			if err := expandSectionsTree(child, dirPath, headerLevel+1); err != nil {
				return err
			}
		}
		return nil
	}

	// Create only leaf sections
	filePath := filepath.Join(parentPath, name+".md")
	consideredSections = append(consideredSections, filePath)

	// In case the file already exists
	// Specific for update command
	if _, err := os.Stat(filePath); err == nil {
		return nil
	}

	// Simple content to start the file
	// The header level matches the level of depth of the section within the section tree
	content := fmt.Sprintf("%s %s\n\n", strings.Repeat("#", headerLevel), section.Title)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}

// sanitizeString enforces snake_case removing characters that make
// your version control system cry
func sanitizeString(path string) string {
	// Unicode -> ASCII
	t := norm.NFD.String(path)
	t = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, t)

	// Regexes magic crap
	nameStyleRegex := regexp.MustCompile(`[^a-zA-Z0-9_./]`)
	singleUnderscoreRegex := regexp.MustCompile(`_+`)

	// Guarantee snake_case
	t = strings.ReplaceAll(t, " ", "_")
	t = strings.ReplaceAll(t, "-", "_")
	t = nameStyleRegex.ReplaceAllString(t, "")
	t = strings.ToLower(t)
	t = singleUnderscoreRegex.ReplaceAllString(t, "_")
	t = strings.ReplaceAll(t, "_.md", ".md")

	return filepath.Clean(t)
}

func deleteNonPresentSections() error {
	toBeKept := make(map[string]struct{}, len(consideredSections))
	for _, s := range consideredSections {
		toBeKept[s] = struct{}{}
	}

	err := filepath.WalkDir(SECTIONS_FOLDER_NAME, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if _, ok := toBeKept[path]; !ok {
			if rmErr := os.Remove(path); rmErr != nil {
				return rmErr
			}
		}
		return nil
	})
	return err
}

func deleteEmptyDirectories() error {
	return filepath.WalkDir(SECTIONS_FOLDER_NAME, func(path string, d os.DirEntry, err error) error {

		fmt.Println(path)

		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		entries, readErr := os.ReadDir(path)
		if readErr != nil {
			return readErr
		}

		if len(entries) == 0 {
			fmt.Println("deleting on second call:", path)
			if rmErr := os.Remove(path); rmErr != nil {
				return rmErr
			}
			return fs.SkipDir // prevents trying to walk on chidren files of removed dir
		}

		return nil
	})
}

// Internally calls the deleteNonPresentSections and deleteEmptyDirectories functions.
// Ensures that the updated sections folder remains clean.
func CleanDir() error {

	err := deleteNonPresentSections()
	if err != nil {
		return err
	}

	return deleteEmptyDirectories()
}
