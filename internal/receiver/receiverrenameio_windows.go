//go:build windows

package receiver

import (
	"os"
	"path/filepath"
)

type pendingFile struct {
	fn string
	f  *os.File
}

func newPendingFile(root *os.Root, fn string) (*pendingFile, error) {
	abs := filepath.Join(root.Name(), fn)
	f, err := os.CreateTemp(filepath.Dir(abs), "temp-rsync-*")
	if err != nil {
		return nil, err
	}
	return &pendingFile{
		fn: abs,
		f:  f,
	}, nil
}

func (p *pendingFile) Name() string {
	return p.fn
}

func (p *pendingFile) Write(buf []byte) (n int, _ error) {
	return p.f.Write(buf)
}

func (p *pendingFile) CloseAtomicallyReplace() error {
	if err := p.f.Close(); err != nil {
		return err
	}
	if err := os.Rename(p.f.Name(), p.fn); err != nil {
		return err
	}
	return nil
}

func (p *pendingFile) Cleanup() error {
	tmpName := p.f.Name()
	err := p.f.Close()
	if err := os.Remove(tmpName); err != nil {
		return err
	}
	return err
}
