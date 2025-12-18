package lens

import (
	"context"
	"os"

	"statelensfs/vfs"
	"statelensfs/state"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type MemDir struct{
	summary string
}

func NewMemDir() *MemDir {
	return &MemDir{
		summary: "summary",
	}
}

func (d *MemDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	return nil
}

func (d *MemDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == d.summary {
		return vfs.NewDynamicFile(state.MemSummary), nil
	}
	return nil, fuse.ENOENT
}

func (d *MemDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		{Name: d.summary, Type: fuse.DT_File},
	}, nil
}
