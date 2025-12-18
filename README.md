# StateLensFS

StateLensFS is a **read-only, userspace filesystem** that exposes live system and runtime state as files.

It is designed to feel like `/proc` or `/sys`, but implemented entirely in userspace using FUSE.  
Consumers interact with it using standard filesystem tools (`cat`, `ls`, `grep`) — no agents, no RPC, no SDKs.

## Why it exists

StateLensFS is intended for control-plane introspection, not high-frequency data paths.

## Motivation

Modern systems often expose runtime state via:
- HTTP APIs
- agents
- metrics endpoints
- bespoke libraries

StateLensFS explores a different approach:

> **What if runtime state was just files?**

By mapping live system information into a filesystem:
- shell scripts can reason about state naturally
- schedulers and supervisors avoid tight coupling
- observability becomes composable
- tooling stays Unix-native

## Features

- Read-only filesystem
- Dynamic file contents (evaluated on read)
- Zero persistence
- Zero configuration
- Minimal dependencies
- Designed for:
  - custom VMs
  - embedded Linux
  - runtime / platform systems
  - schedulers and supervisors

## Filesystem Layout

```
/statelens
├── cpu
│   └── summary
├── mem
│   └── summary
└── net
    ├── interfaces
    └── routes
```

## Examples

```
$ cat /statelens/cpu/summary
 18:42:31 up 3 days,  2 users,  load average: 0.15, 0.09, 0.05

$ cat /statelens/mem/summary
              total        used        free      shared  buff/cache   available
Mem:           3.8G        512M        2.4G         12M        896M        3.1G
```

## Architecture

StateLensFS is structured into four layers:

```
mount/    → FUSE filesystem wiring
lens/     → directory and file semantics
state/    → collectors (system state providers)
vfs/      → generic filesystem primitives
```

- Collectors gather live data (e.g. uptime, free, ip)
- Lenses map collectors into directories and files
- Filesystem exposes everything via FUSE

This separation makes it easy to add new lenses without touching mount logic.

## Building

StateLensFS is a single static binary.

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o statelensfs
```

> Note: StateLensFS requires FUSE (fuse3) to be available on the host system.

## Running (Development)

```
sudo mkdir -p /mnt/statelens
sudo ./statelensfs /mnt/statelens
```

Then in another shell:

```
ls /mnt/statelens
cat /mnt/statelens/cpu/summary
```

## Running in a Custom VM / Init System

StateLensFS is designed to be mounted at boot, similar to /proc.

## Example init snippet

```
mkdir -p /statelens
/sbin/statelensfs /statelens &
```

Once mounted, the filesystem is immediately available to all processes.

## Use Cases

- Runtime introspection for schedulers
- Embedded system state exposure
- Lightweight observability without agents
- Shell-driven automation
- Platform debugging surfaces

## Extending StateLensFS

Adding a new lens typically involves:

- Writing a collector in state/
- Mapping it to files in lens/
- Wiring it into the root directory

No persistence or schema required.

## Project Status

StateLensFS is intentionally minimal.

It focuses on:

- correctness
- clear filesystem semantics
- composable, Unix-friendly interfaces

The project favors simplicity and debuggability over feature breadth.
New lenses and interfaces may evolve as use cases emerge.

## License

MIT