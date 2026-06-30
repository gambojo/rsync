//go:build windows

package rsynctest

import "testing"

func CreateDummyDeviceFiles(t *testing.T, dir string) {
	t.Skip("Windows does not support special files (devices/FIFOs/sockets)")
}

func VerifyDummyDeviceFiles(t *testing.T, source, dest string) {
	t.Skip("Windows does not support special files (devices/FIFOs/sockets)")
}
