package tools

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/phaus/go-unix-wrapper/sys"
)

func (mounter *Mounter) Mount() (string, error) {
	err := checkMountpoint(mounter.mountPoint)
	if err != nil {
		return "", err
	}
	sudoBin, err := sys.GetPath("sudo")
	if err != nil {
		return "", err
	}
	mountBin, err := sys.GetPath("mount")
	if err != nil {
		return "", err
	}
	cmdCmd := exec.Command(sudoBin, mountBin, "-t", mounter.filesystem, mounter.device, mounter.mountPoint)
	log.Printf("running cmd %v", cmdCmd)
	cmdResult, cmdErr := sys.RunCmd(cmdCmd)
	if cmdErr != nil {
		return "", cmdErr
	}
	return cmdResult, nil
}

func (mounter *Mounter) Umount() (string, error) {
	sudoBin, err := sys.GetPath("sudo")
	if err != nil {
		return "", err
	}
	diskutilBin, err := sys.GetPath("diskutil")
	if err != nil {
		return "", err
	}
	cmdCmd := exec.Command(sudoBin, diskutilBin, "unmount", mounter.mountPoint)
	log.Printf("running cmd %v", cmdCmd)
	cmdResult, cmdErr := sys.RunCmd(cmdCmd)
	if cmdErr != nil {
		return "", cmdErr
	}
	err = mounter.cleanupMountpoint(mounter.mountPoint)
	if err != nil {
		return "", err
	}
	return cmdResult, nil
}

func (mounter *Mounter) ListDevices() (string, error) {
	diskutilBin, err := sys.GetPath("diskutil")
	if err != nil {
		return "", err
	}
	cmdCmd := exec.Command(diskutilBin, "list")
	log.Printf("running cmd %v", cmdCmd)
	cmdResult, cmdErr := sys.RunCmd(cmdCmd)
	if cmdErr != nil {
		return "", cmdErr
	}
	return cmdResult, nil
}

func checkMountpoint(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("%s does not exists. Creating it!\n", path)
		sudoBin, err := sys.GetPath("sudo")
		if err != nil {
			return err
		}
		cmdCmd := exec.Command(sudoBin, "mkdir", "-p", path)
		log.Printf("running cmd %v", cmdCmd)
		cmdResult, cmdErr := sys.RunCmd(cmdCmd)
		if cmdErr != nil {
			return cmdErr
		}
		log.Printf("Created %s: %s\n", path, cmdResult)
	}
	return nil
}

func (mounter *Mounter) cleanupMountpoint(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("%s does not exists. Existing!\n", path)
		return nil
	}
	empty, err := mounter.isEmpty(path)
	if err != nil {
		return err
	}
	if !empty {
		return fmt.Errorf("%s not empty, cannot delete", path)
	} else {
		sudoBin, err := sys.GetPath("sudo")
		if err != nil {
			return err
		}
		cmdCmd := exec.Command(sudoBin, "rm", "-rf", path)
		log.Printf("running cmd %v", cmdCmd)
		cmdResult, cmdErr := sys.RunCmd(cmdCmd)
		if cmdErr != nil {
			return cmdErr
		}
		log.Printf("Created %s: %s\n", path, cmdResult)
	}
	return nil
}
