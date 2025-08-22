// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package cmd

import (
	"fmt"
	"mdm/application"
	"mdm/constants"
	"mdm/errors"
	"mdm/log"

	"github.com/spf13/cobra"
)

const (
	SUCCESS_MESSAGE    = "Successfully generated sections"
	SAMPLE_SCHEMA_FLAG = "sample"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Generate the '%s' folder corresponding to the schema file", constants.SECTIONS_FOLDER_NAME),
	Long:  fmt.Sprintf("Reads the contents of the '%s' file on the current path and generates the corresponding document structure.", constants.SCHEMA_FILE_NAME),
	Run: func(cmd *cobra.Command, args []string) {
		if application.SectionsFolderExists() {
			log.FatalError(errors.ExistingFolderError(constants.SECTIONS_FOLDER_NAME))
		}

		generateSampleSchema, _ := cmd.Flags().GetBool(SAMPLE_SCHEMA_FLAG)
		schemaFile := application.SchemaFileExists()

		if generateSampleSchema {
			if schemaFile {
				log.FatalError(errors.ExistingFileError(constants.SCHEMA_FILE_NAME))
			}

			err := application.GenerateSampleSchema()
			if err != nil {
				log.FatalError(err)
			}
		}

		sections, err := application.LoadSchema()
		if err != nil {
			log.Warning(errors.SchemaFileNotFound())
		}

		err = application.GenerateDocumentSections(sections)
		if err != nil {
			log.FatalError(err)
		}

		log.Success(SUCCESS_MESSAGE)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP(SAMPLE_SCHEMA_FLAG, "s", false, "generate a sample 'schema.yaml' file")
}
