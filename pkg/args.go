package pkg

import "fmt"

type ProgramArguments struct {
	BlockSize    uint   `arg:"-b, --block-size" help:"block size as kb" default:"1024"`
	Count        uint   `arg:"-c, --count" help:"iteration count of block size. If count is 1024 and block size is 1024kb, data size will be 1024 x 1024kb = 1GB" default:"1024"`
	TestFilePath string `arg:"-f, --file-path" help:"test file to run benchmark on. If you want to benchmark another disk you can specify mount point on that" default:"./disk-write-test-file"`
}

func (p *ProgramArguments) String() string {
	return fmt.Sprintf("Block Size: %dkb Count: %d File Path: %s", p.BlockSize, p.Count, p.TestFilePath)
}
