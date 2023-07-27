package utils

import "syscall"

func GetLinuxKernelVersion() (int, int, int) {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return 0, 0, 0
	}

	var (
		major = int(uname.Release[0] - '0')
		minor = int(uname.Release[2] - '0')
		patch = int(uname.Release[4] - '0')
	)

	return major, minor, patch
}
