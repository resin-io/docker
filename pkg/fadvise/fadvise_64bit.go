// +build linux,amd64 linux,arm64 linux,s390x

package fadvise

import "syscall"

func posixFadvise64(fd int, offset, length int64, advice int) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_FADVISE64, uintptr(fd), uintptr(offset), uintptr(offset>>32), uintptr(length), uintptr(length>>32), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}
