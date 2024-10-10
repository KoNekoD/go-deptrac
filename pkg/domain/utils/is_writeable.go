package utils

import "golang.org/x/sys/unix"

func IsWriteable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}
