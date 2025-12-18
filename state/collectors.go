package state

import (
	"bytes"
	"os/exec"
)

func run(cmd string, args ...string) ([]byte, error) {
	c := exec.Command(cmd, args...)
	var out bytes.Buffer
	c.Stdout = &out
	c.Stderr = &out
	err := c.Run()
	return out.Bytes(), err
}

func CPUSummary() ([]byte, error) {
	return run("uptime")
}

func MemSummary() ([]byte, error) {
	return run("free", "-h")
}

func NetInterfaces() ([]byte, error) {
	return run("ip", "a")
}

func NetRoutes() ([]byte, error) {
	return run("ip", "r")
}
