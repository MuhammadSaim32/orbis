/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MuhammadSaim32/orbis/internals/devices"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var deviceCmd = &cobra.Command{
	Use: "device",
	Run: func(cmd *cobra.Command, args []string) {
		devices.ListDisk()
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
