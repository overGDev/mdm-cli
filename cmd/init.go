// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package cmd

import (
	"fmt"
	"mdm/internal/apperrors"
	"mdm/internal/application"
	"mdm/internal/applog"

	"github.com/spf13/cobra"
)

const (
	SAMPLE_SCHEMA_FLAG = "sample"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Generate the '%s' folder corresponding to the schema file", application.SECTIONS_FOLDER_NAME),
	Long:  fmt.Sprintf("Reads the contents of the '%s' file on the current path and generates the corresponding document structure.", application.SCHEMA_FILE_NAME),
	Run: func(cmd *cobra.Command, args []string) {
		if application.SectionsFolderExists() {
			applog.FatalError(apperrors.ExistingFolderError(application.SECTIONS_FOLDER_NAME))
		}

		generateSampleSchema, _ := cmd.Flags().GetBool(SAMPLE_SCHEMA_FLAG)
		schemaFile := application.SchemaFileExists()

		if generateSampleSchema {
			if schemaFile {
				applog.FatalError(apperrors.ExistingFileError(application.SCHEMA_FILE_NAME))
			}

			err := application.GenerateSampleSchema()
			if err != nil {
				applog.FatalError(err)
			}
		}

		sections, err := application.LoadSchema()
		if err != nil {
			applog.Warning(apperrors.SchemaFileNotFound())
		}

		err = application.GenerateDocumentSections(sections)
		if err != nil {
			applog.FatalError(err)
		}

		err = application.GenerateGitHubAction()
		if err != nil {
			applog.FatalError(err)
		}

		err = application.GenerateGitIgnore()
		if err != nil {
			applog.FatalError(err)
		}

		err = application.GenerateScript()
		if err != nil {
			applog.FatalError(err)
		}

		applog.Success("Successfully generated sections")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP(SAMPLE_SCHEMA_FLAG, "s", false, "generate a sample 'schema.yaml' file")
}
