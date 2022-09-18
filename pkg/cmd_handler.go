package pkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dustin/go-humanize"
)

func Run(args *ProgramArguments) (string, error) {
	defer func() {
		if err := os.Remove(args.TestFilePath); err != nil {
			log.Printf("Failed to remove test file %s because of %s. Please manually remove it.", args.TestFilePath, err.Error())
		}
	}()
	start := time.Now()

	cmd := exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", args.TestFilePath), "conv=sync", fmt.Sprintf("bs=%dk", args.BlockSize), fmt.Sprintf("count=%d", args.Count))
	stderr, _ := cmd.StderrPipe()
	defer func() {
		_ = stderr.Close()
	}()

	if err := cmd.Start(); err != nil {
		return "", err
	}

	printStdErrResult(stderr)

	since := time.Since(start)

	speed := calculateSpeed(getTotalBytes(args), since.Seconds())

	return fmt.Sprintf("Disk Write Speed: %s/s\n", humanize.Bytes(speed)), nil
}

func printStdErrResult(stderr io.ReadCloser) {
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func getTotalBytes(args *ProgramArguments) uint64 {
	return uint64((args.BlockSize * 1000) * args.Count)
}

func calculateSpeed(bytes uint64, sinceAsSeconds float64) uint64 {
	return uint64(float64(bytes) / sinceAsSeconds)
}
