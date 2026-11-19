package main

import (
	"fmt"
)

func help(Arguments []string) {
	fmt.Println("Orbis — Linux Disk Management Tool")
	fmt.Println("==================================")
	fmt.Println()
	fmt.Println("View information about block storage devices on your Linux system.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  orbis <command>")
	fmt.Println()
	fmt.Println("COMMANDS")
	fmt.Println("  listDevices    List block devices with their name, vendor, and model")
	fmt.Println("  help           Display this help message")
	fmt.Println()
	fmt.Println("EXAMPLES")
	fmt.Println("  orbis listDevices")
	fmt.Println("  orbis help")
	fmt.Println()
	fmt.Println("DEVELOPMENT (Go)")
	fmt.Println("  go run . listDevices")
	fmt.Println("  go run . help")
}
