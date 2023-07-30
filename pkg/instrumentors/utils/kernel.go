package utils

import (
	"fmt"
	"log"
	"syscall"

	"github.com/hashicorp/go-version"
)

func GetLinuxKernelVersion() (*version.Version, error) {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return nil, err
	}

	major := int(uname.Release[0] - '0')
	minor := int(uname.Release[2] - '0')
	patch := int(uname.Release[4] - '0')
	log.Printf("############ Kernel version is: %d.%d%d", major, minor, patch)
	versionStr := fmt.Sprintf("%d.%d%d", major, minor, patch)

	ver, err := version.NewVersion(versionStr)
	if err != nil {
		return nil, err
	}

	return ver, nil
}
