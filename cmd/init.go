// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package cmd

import (
	"fmt"
	"mdm/constants"
	"mdm/errors"
	"mdm/log"
	"mdm/validation"

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
		if validation.SectionsFolderExists() {
			log.FatalError(errors.ExistingFolderError(constants.SECTIONS_FOLDER_NAME))
		}

		generateSampleSchema, _ := cmd.Flags().GetBool(SAMPLE_SCHEMA_FLAG)
		schemaFile := validation.SchemaFileExists()

		if generateSampleSchema {
			if schemaFile {
				log.FatalError(errors.ExistingFileError(constants.SCHEMA_FILE_NAME))
			}

			err := validation.GenerateSampleSchema()
			if err != nil {
				log.FatalError(err)
			}
		}

		// Generate the sections folder
		if !schemaFile {
			log.FatalError(errors.SchemaFileNotFound())
		}

		log.Success(SUCCESS_MESSAGE)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP(SAMPLE_SCHEMA_FLAG, "s", false, "generate a sample 'schema.yaml' file")
}
