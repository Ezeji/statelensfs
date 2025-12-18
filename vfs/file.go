package vfs

import (
	"context"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type DynamicFile struct {
	read func() ([]byte, error)
}

func NewDynamicFile(read func() ([]byte, error)) *DynamicFile {
	return &DynamicFile{read: read}
}

func (f *DynamicFile) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = 0444
	a.Size = 4096 
	return nil
}

func (f *DynamicFile) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	return f, nil
}

func (f *DynamicFile) ReadAll(ctx context.Context) ([]byte, error) {
	return f.read()
}
