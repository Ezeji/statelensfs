package main

import (
	"log"
	"os"

	"statelensfs/mount"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: statelensfs MOUNTPOINT")
		os.Exit(1)
	}

	mountpoint := os.Args[1]

	c, err := fuse.Mount(
		mountpoint,
		fuse.ReadOnly(),
		fuse.FSName("statelens"),
		fuse.Subtype("statelensfs"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if err := fs.Serve(c, &mount.StateFS{}); err != nil {
		log.Fatal(err)
	}
}
