package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dustin/go-humanize"
)

func Run(args *ProgramArguments) (string, error) {
	removeTestFile(args)
	defer func() {
		removeTestFile(args)
	}()
	fmt.Printf("Starting with %s\n", args.String())

	_ = exec.Command("echo", "3 > /proc/sys/vm/drop_caches").Run()

	start := time.Now()
	cmd := exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", args.TestFilePath), "conv=fsync", fmt.Sprintf("bs=%dk", args.BlockSize), fmt.Sprintf("count=%d", args.Count))
	_, err := cmd.Output()
	if err != nil {
		return "", err
	}

	since := time.Since(start)

	totalBytes := getTotalBytes(args)
	speed := calculateSpeed(totalBytes, since.Seconds())

	return fmt.Sprintf("Disk Write Speed: %s/s Elapsed Time: %fs Total Written Data Size: %s\n", humanize.Bytes(speed), since.Seconds(), humanize.Bytes(totalBytes)), nil
}

func getTotalBytes(args *ProgramArguments) uint64 {
	return uint64((args.BlockSize * 1000) * args.Count)
}

func calculateSpeed(bytes uint64, sinceAsSeconds float64) uint64 {
	return uint64(float64(bytes) / sinceAsSeconds)
}

func removeTestFile(args *ProgramArguments) {
	if _, err := os.Stat(args.TestFilePath); !errors.Is(err, os.ErrNotExist) {
		if err := os.Remove(args.TestFilePath); err != nil {
			log.Printf("Failed to remove test file %s because of %s. Please manually remove it.", args.TestFilePath, err.Error())
		}
	}
}
