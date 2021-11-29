package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

var dir string

func init() {
	flag.StringVar(&dir, "d", "./", "检查的项目目录")
	flag.Parse()
}

func main() {
	govet := exec.Command("go", "vet")
	golint := exec.Command("golangci-lint", "run", "./...")
	govet.Dir = dir
	golint.Dir = dir
	fmt.Println("== go vet ==")
	_ = execute(govet)
	fmt.Println("== golangci-lint ==")
	_ = execute(golint)
}

func execute(cmd *exec.Cmd) error {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %s", err.Error())
		return err
	}

	go func() {
		_ = asyncLog(stdout)
	}()
	go func() {
		_ = asyncLog(stderr)
	}()

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s", err.Error())
		return err
	}

	return nil
}

func asyncLog(reader io.ReadCloser) error {
	var cache string
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				err = nil
			}
			return err
		}
		if num > 0 {
			oByte := buf[:num]
			oSlice := strings.Split(string(oByte), "\n")
			line := strings.Join(oSlice[:len(oSlice)-1], "\n")
			fmt.Printf("%s%s\n", cache, line)
			cache = oSlice[len(oSlice)-1]
		}
	}
}
