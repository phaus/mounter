package tools

import (
	"io"
	"os"
)

type Mounter struct {
	mountPoint string
	device     string
	filesystem string
}

const DEFAULT_FS = "msdos"

func NewMounter(device, mountPoint, filesystem string) *Mounter {
	if filesystem == "" {
		filesystem = DEFAULT_FS
	}
	return &Mounter{
		mountPoint: mountPoint,
		device:     device,
		filesystem: filesystem,
	}
}

func (mounter *Mounter) isEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
