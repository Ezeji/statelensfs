package lens

import (
	"context"
	"os"

	"statelensfs/vfs"
	"statelensfs/state"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type NetDir struct{
	interfaces string
	routes string
}

func NewNetDir() *NetDir {
	return &NetDir{
		interfaces: "interfaces",
		routes:     "routes",
	}
}

func (d *NetDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	return nil
}

func (d *NetDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	switch name {
	case d.interfaces:
		return vfs.NewDynamicFile(state.NetInterfaces), nil
	case d.routes:
		return vfs.NewDynamicFile(state.NetRoutes), nil
	}
	return nil, fuse.ENOENT
}

func (d *NetDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		{Name: d.interfaces, Type: fuse.DT_File},
		{Name: d.routes, Type: fuse.DT_File},
	}, nil
}
