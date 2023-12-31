package Myls

import (
	"os"
	"time"
)

type File struct {
	Info        os.DirEntry
	Permissions os.FileMode // Adding Permissions field
	Links       uint64      // Adding Links field
	Owner       string      // Adding Owner field
	Group       string      // Adding Group field
	Size        int64       // Adding Size field
	ModTime     time.Time   // Adding ModTime field
	blockSize   int64       // Adding blockSize field
}

type Flags struct {
	L  bool
	R  bool
	A  bool
	Rr bool
	T  bool
}

var (
	Fail        []string
	Success     []string
	TotalBlocks int64
)

const (
	Blue    = "\033[1;34m"
	Green   = "\033[0;32m"
	Cyan    = "\033[0;36m"
	Yellow  = "\033[30;43m" // black foreground, yellow background
	Magenta = "\033[0;35m"
	Red     = "\033[0;31m"
	BRed    = "\033[30;41m" // black foreground, red background
	Reset   = "\033[0m"
)
