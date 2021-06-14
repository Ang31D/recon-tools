package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"flag"
)

func main() {
	flag.Parse()
	domains := getInput()

	for domains.Scan() {
		domain := domains.Text()
		fmt.Println(reverseDomain(domain))
	}
	if err := domains.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "revasset", err)
	}
}

// https://gist.github.com/nesv/06175c88bc37e56bc739d02b5fe284f5
func getInput() (*bufio.Scanner) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return bufio.NewScanner(os.Stdin)
	}
	return bufio.NewScanner(strings.NewReader(flag.Arg(0)))
}

func reverseDomain(domain string) (string) {
	parts := strings.Split(domain, ".")
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}
func init() {
	flag.Usage = func() {
		h := "Reverse Asset (GOLANG) - Reverse each '.' position (ex. com.google / 1.0.0.127).\n"
		h += "supports stdin/stdout for workflow integration\n\n"
		h += "Usage: \n"
		h += "  revasset (<domain>)\n\n"

		h += "Examples:\n"
		h += "  revasset example.com\n"
		h += "  cat domains.txt | revasset\n"

		fmt.Fprint(os.Stderr, h)
	}
}
