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

func showlast(s string, lines int) {
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
	lastbuf := make([]byte, bsize)
	last := time.Now()

	for {
		n, err := io.ReadAtLeast(os.Stdin, buf, bsize)
		if err != nil || n < bsize {
			showlast(string(lastbuf)+string(buf[0:n]), lines)
			break
		}
		if time.Since(last) >= time.Duration(interval)*time.Second {
			showlast(string(lastbuf)+string(buf), lines)
			last = time.Now()
		}
		buf, lastbuf = lastbuf, buf
	}
}

func main() {
	var lines, interval int
	flag.IntVar(&lines, "n", 20, "Number of lines to output")
	flag.IntVar(&interval, "interval", 1, "Delay between output")
	flag.Parse()
	pipetail(lines, interval)
}
