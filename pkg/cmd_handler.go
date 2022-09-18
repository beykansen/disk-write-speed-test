package pkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

	stderrResults := getAndPrintStdErrResult(stderr)

	since := time.Since(start)

	bytes, err := getTotalBytes(stderrResults)
	if err != nil {
		return "", err
	}

	speed := calculateSpeed(bytes, since.Seconds())

	return fmt.Sprintf("Disk Write Speed: %s/s\n", humanize.Bytes(speed)), nil
}

func getAndPrintStdErrResult(stderr io.ReadCloser) []string {
	stderrResults := make([]string, 0)
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		stderrResults = append(stderrResults, scanner.Text())
	}
	return stderrResults
}

func getTotalBytes(results []string) (uint64, error) {
	actualResult := ""
	for _, k := range results {
		if strings.Contains(k, "transferred in") {
			actualResult = k
		}
	}
	bytesAsStr := strings.TrimSpace(strings.Split(actualResult, "bytes ")[0])
	bytes, err := strconv.ParseUint(bytesAsStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return bytes, nil
}

func calculateSpeed(bytes uint64, sinceAsSeconds float64) uint64 {
	return uint64(float64(bytes) / sinceAsSeconds)
}
