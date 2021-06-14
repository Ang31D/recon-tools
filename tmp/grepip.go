package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var option struct {
	match_contains bool
	extract_ipv4 bool
	invert_match bool
}

func main() {
	flag.Parse()
	data := getInput()

	fmt.Fprintf(os.Stdout, "option.match_contains): %t\n", option.match_contains)
	fmt.Fprintf(os.Stdout, "option.invert_match): %t\n", option.invert_match)

	for data.Scan() {
		value := data.Text()
		// match on IPv4
		if !(option.invert_match) {
//			fmt.Println("inside ! invert_match")
			// exact match
			if (! (option.match_contains)) {
//				fmt.Println("inside exact_match")
				if isIPv4(value) {
					fmt.Fprintf(os.Stdout, "isIPv4: '%s'\n", value)
				}
			} else {
				if (containsIPv4(value)) {
					fmt.Fprintf(os.Stdout, "containsIPv4: '%s'\n", value)
				}
			}
		} else {
			// exact match
			if (! (option.match_contains)) {
				if (! (isIPv4(value))) {
					fmt.Fprintf(os.Stdout, "NOT isIPv4: '%s'\n", value)
				}
			} else {
				if (! (containsIPv4(value))) {
					fmt.Fprintf(os.Stdout, "NOT containsIPv4: '%s'\n", value)
				}
			}
		}
	}
}
func isIPv4(value string) bool {
//	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
//	pattern := "^" + numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock + "$"
	pattern := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"

	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}
func containsIPv4(value string) bool {
//	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
//	pattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
	pattern := "(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])"

	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func getInput() *bufio.Scanner {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return bufio.NewScanner(os.Stdin)
	}
	return bufio.NewScanner(strings.NewReader(flag.Arg(0)))
}

func init() {
	flag.BoolVar(&option.match_contains, "g", true, "")
	flag.BoolVar(&option.match_contains, "match-on", false, "")

	flag.BoolVar(&option.extract_ipv4, "o", false, "")
	flag.BoolVar(&option.extract_ipv4, "extract-ipv4", false, "")

	flag.BoolVar(&option.invert_match, "v", false, "")
	flag.BoolVar(&option.invert_match, "invert-match", false, "")

	flag.Usage = func() {
		h := "Grep IP (GOLANG) - Extract / Exclude IPv4 addresses\n"
		h += "supports stdin/stdout for workflow integration\n\n"

		h += "Usage:\n"
		h += "  grepip (<string>) [options]\n\n"

		h += "Options:\n"
		h += "  -g, --match-on      match on IPv4, otherwise exact\n\n"
		//		h += "  -o, --extract       extract IPv4 match\n\n"
		h += "  -v, --invert-match  exclude IPv4 match\n"

		h += "Examples:\n"
		h += "  grepip 'string 127.0.0.1#' [options]\n"
		h += "  cat data.txt | grepip [options]\n"

		fmt.Fprint(os.Stderr, h)
	}

}
