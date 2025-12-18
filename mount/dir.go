package mount

import (
	"context"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type RootDir struct{
	cpu fs.Node
	mem fs.Node
	net fs.Node
}

func (d *RootDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	return nil
}

func (d *RootDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	switch name {
	case "cpu":
		return d.cpu, nil
	case "mem":
		return d.mem, nil
	case "net":
		return d.net, nil
	}
	return nil, fuse.ENOENT
}

func (d *RootDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		{Name: "cpu", Type: fuse.DT_Dir},
		{Name: "mem", Type: fuse.DT_Dir},
		{Name: "net", Type: fuse.DT_Dir},
	}, nil
}
