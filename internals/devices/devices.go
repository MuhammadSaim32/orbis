package devices

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ListDisk() {

	type DeviceInfo struct {
		Name   string
		Vendor string
		Model  string
		Size   int
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

		size, err := os.ReadFile(fullPath + "/size")

		if err != nil {
			fmt.Println("Error reading block devices:", err)
			return
		}

		sectors, err := strconv.Atoi(strings.TrimSpace(string(size)))

		if err != nil {
			return
		}

		diskInfoSlice = append(diskInfoSlice, DeviceInfo{
			Name:   item.Name(),
			Vendor: strings.TrimSpace(string(vender)),
			Model:  strings.TrimSpace(string(model)),
			Size:   (sectors * 512) / (1024 * 1024 * 1024),
		})
	}

	for _, item := range diskInfoSlice {

		fmt.Printf("Name : %s | Vender : %s | Model :%s | Size : %v GB \n", item.Name, item.Vendor, item.Model, item.Size)
	}

}
