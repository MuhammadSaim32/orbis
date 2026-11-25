/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MuhammadSaim32/orbis/internals/devices"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "List available storage devices",
	Long:  "List information about available Linux block devices, including their names, vendors, models, and sizes.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		devices.ListDisk()
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)

}
