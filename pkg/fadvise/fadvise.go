package fadvise

import "os"

func PosixFadvise(file *os.File, offset, length int64, advice int) error {
	return posixFadvise64(int(file.Fd()), offset, length, advice)
}
