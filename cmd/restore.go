/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/richbai90/xfer-bin/pkg/action"
)

var source, dest string;

// initCmd represents the init command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore an archive",
	Run: action.Restore(&source, &dest),
}

func init() {
	restoreCmd.Flags().StringVarP(&source, "source", "s", "", "Specify the file or folder path to restore from")
	restoreCmd.Flags().StringVarP(&dest, "dest", "d", "", "Specify the destination directory")
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
