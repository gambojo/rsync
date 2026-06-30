//go:build windows

package rsynctest

import (
	"os"
	"testing"
)

func CreateDummyDeviceFiles(t *testing.T, dir string) {
	t.Skip("Windows does not support special files (devices/FIFOs/sockets)")
}

func VerifyDummyDeviceFiles(t *testing.T, source, dest string) {
	t.Skip("Windows does not support special files (devices/FIFOs/sockets)")
}

// StatUidGid is not meaningful on Windows (no Unix file ownership).
//
// Callers are supposed to check os.Getuid()==-1 and skip.
// Hence, this function must compile but is never reached.
func StatUidGid(t *testing.T, fi os.FileInfo) (uid, gid int) {
	t.Fatal("BUG: StatUidGid not skipped on Windows")
	return 0, 0
}
