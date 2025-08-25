/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mdm/internal/apperrors"
	"mdm/internal/application"
	"mdm/internal/applog"
	"slices"

	"github.com/spf13/cobra"
)

var validArgs = []string{"sections"}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update repository content",
	Args:  cobra.ExactArgs(1),
	Long:  `Update your repository folders and files according to changes on the configuration files`,
	Run: func(cmd *cobra.Command, args []string) {
		if !slices.Contains(validArgs, args[0]) {
			applog.Warning(apperrors.InvalidArgument("update"))
		}

		if !application.SchemaFileExists() {
			applog.FatalError(apperrors.MissingFileError(application.SCHEMA_FILE_NAME))
		}

		sections, err := application.LoadSchema()
		if err != nil {
			applog.Warning(apperrors.SchemaFileNotFound())
		}

		err = application.GenerateDocumentSections(sections)
		if err != nil {
			applog.FatalError(err)
		}

		deleteNonPresent, _ := cmd.Flags().GetBool("delete-non-present")

		if deleteNonPresent {
			err := application.CleanDir()
			if err != nil {
				applog.FatalError(err)
			}
		}

		applog.Success(fmt.Sprintf("successfully updated %s", application.SECTIONS_FOLDER_NAME))

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolP("delete-non-present", "d", false, fmt.Sprintf("deletes all the sections not present on %s", application.SCHEMA_FILE_NAME))
}
