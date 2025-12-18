package mount

import (
	"statelensfs/lens"

	"bazil.org/fuse/fs"
)

type StateFS struct{}

func (StateFS) Root() (fs.Node, error) {
	return &RootDir{
		cpu: lens.NewCPUDir(),
		mem: lens.NewMemDir(),
		net: lens.NewNetDir(),
	}, nil
}
