package lens

import (
	"context"
	"os"

	"statelensfs/vfs"
	"statelensfs/state"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type CPUDir struct{
	summary string
}

func NewCPUDir() *CPUDir {
	return &CPUDir{
		summary: "summary",
	}
}

func (d *CPUDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	return nil
}

func (d *CPUDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == d.summary {
		return vfs.NewDynamicFile(state.CPUSummary), nil
	}
	return nil, fuse.ENOENT
}

func (d *CPUDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		{Name: d.summary, Type: fuse.DT_File},
	}, nil
}
