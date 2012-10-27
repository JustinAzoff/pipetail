package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const LINESIZE = 2048

func showlast(buffer []byte, lines int) {
	s := string(buffer)
	parts := strings.Split(s, "\n")
	num_parts := len(parts)
	first := num_parts - lines - 1
	last := num_parts - 1
	for _, line := range parts[first:last] {
		fmt.Println(line)
	}
}

func pipetail(lines, interval int) {
	bsize := lines * LINESIZE
	buf := make([]byte, bsize)
	last := time.Now()

	for {
		_, err := io.ReadAtLeast(os.Stdin, buf, bsize)
		if err != nil {
			break
		}
		if time.Since(last) >= time.Duration(interval)*time.Second {
			showlast(buf, lines)
			last = time.Now()
		}
	}
    showlast(buf, lines)
}

func main() {
	var lines, interval int
	flag.IntVar(&lines, "n", 20, "Number of lines to output")
	flag.IntVar(&interval, "interval", 1, "Delay between output")
	flag.Parse()
	pipetail(lines, interval)
}
