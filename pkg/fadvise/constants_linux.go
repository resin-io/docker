// +build !s390x

package fadvise

const (
	POSIX_FADV_NORMAL     = 0 /* No further special treatment.  */
	POSIX_FADV_RANDOM     = 1 /* Expect random page references.  */
	POSIX_FADV_SEQUENTIAL = 2 /* Expect sequential page references.  */
	POSIX_FADV_WILLNEED   = 3 /* Will need these pages.  */
	POSIX_FADV_DONTNEED   = 4 /* Don't need these pages.  */
	POSIX_FADV_NOREUSE    = 5 /* Data will be accessed once.  */
)
