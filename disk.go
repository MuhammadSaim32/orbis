package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func listDisks(Arguments []string) {
	fmt.Println(Arguments)

	type DeviceInfo struct {
		Name   string
		Vendor string
		Model  string
	}

	const path = "/sys/block"

	data, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading block devices:", err)
		return
	}

	var diskInfoSlice []DeviceInfo

	for _, item := range data {
		fullPath := filepath.Join(path, item.Name())
		vender, err := os.ReadFile(fullPath + "/device/vendor")
		if err != nil {
			fmt.Println("Error reading block devices:", err)
			return
		}
		model, err := os.ReadFile(fullPath + "/device/model")

		if err != nil {
			fmt.Println("Error reading block devices:", err)
			return
		}

		diskInfoSlice = append(diskInfoSlice, DeviceInfo{
			Name:   item.Name(),
			Vendor: strings.TrimSpace(string(vender)),
			Model:  strings.TrimSpace(string(model)),
		})
	}

	for _, item := range diskInfoSlice {

		fmt.Printf("Name : %s | Vender : %s | Model :%s\n", item.Name, item.Vendor, item.Model)
	}

}
