package fadvise

import "syscall"

func posixFadvise64(fd int, offset, length int64, advice int) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_ARM_FADVISE64_64, uintptr(fd), uintptr(advice), uintptr(offset), uintptr(offset>>32), uintptr(length), uintptr(length>>32))
	if e1 != 0 {
		err = e1
	}
	return
}
