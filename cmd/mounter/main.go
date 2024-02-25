package main

import (
	"fmt"
	"os"

	"git-lab.de/philipp/mounter/lib/tools"
)

func main() {
	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		help()
	}

	var output, cmd, device, mountPoint, filesystem string
	var err error
	mode := len(argsWithoutProg)
	if mode == 1 {
		cmd = argsWithoutProg[0]
	}
	if mode == 2 {
		cmd = argsWithoutProg[0]
		mountPoint = argsWithoutProg[1]
	}
	if mode == 3 {
		cmd = argsWithoutProg[0]
		device = argsWithoutProg[1]
		mountPoint = argsWithoutProg[2]
	}

	if mode == 4 {
		cmd = argsWithoutProg[0]
		device = argsWithoutProg[1]
		mountPoint = argsWithoutProg[2]
		filesystem = argsWithoutProg[3]
	}
	mounter := tools.NewMounter(device, mountPoint, filesystem)
	switch cmd {
	case "list":
		output, err = mounter.ListDevices()
	case "mount":
		output, err = mounter.Mount()
	case "umount":
		output, err = mounter.Umount()
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

func help() {
	out := `
	mounter [CMD] ([device] [mount-point] ([filesystem]))
	CMD:
		list - list devices
		mount - mounts a device, requires [device] [mount-point] and - OPTIONAL - [filesystem]
		umount - umounts a device, requires [mount-point] 
	device:
		A unix device (e.g. /dev/disk1s1)
	mount-point:
		A Folder where the device should be mounted (e.g. /Volumes/EFI)
	filesystem:
		A File-System - defaults to 'msdos'
	`
	fmt.Println(out)
	os.Exit(64)
}
