// This file is part of shelldoc.
// © 2023, Mirko Boehm <mirko@kde.org> and the shelldoc contributors
// SPDX-License-Identifier: GPL-3.0

package cmd

import (
	"fmt"

	"github.com/DFlicki/shelldoc/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the shelldoc version",
	Long:  `Print the shelldoc version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
