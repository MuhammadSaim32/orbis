package main

import (
	"fmt"
	"os"
)

func main() {

	// fmt.Println("Orbis - Linux Disk Management Tool\n")

	actions := map[string]func([]string){
		"listDevices": listDisks,
		"help":        help,
	}

	Args := os.Args

	if len(Args) > 1 {
		if fn, ok := actions[Args[1]]; ok {
			fn(Args[2:])
		}
	} else {
		fmt.Printf("NO Argument provided....")
	}

}
