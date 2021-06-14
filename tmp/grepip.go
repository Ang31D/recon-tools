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
	contains_ipv4 bool
	extract_ipv4 bool
	invert_match bool
	ipv4_alike bool
	as_ipv4 bool
	unique bool
	debug bool
}

func main() {
	flag.Parse()

	sc := getInput()

	seen := make(map[string]bool)

	for sc.Scan() {
		val := sc.Text()

		result := parseVal(val)

		//fmt.Println(strings.Join(result, "\n"))
		for _, v := range result {
			if seen[v] {
				continue
			}
			//fmt.Println(v)
			fmt.Fprintf(os.Stdout, "%s\n", v)
			seen[v] = true
		}
	}
}

func parseVal(val string) []string {
	out := []string{}

	// match on IPv4 (or alike)
	if !(option.invert_match) {
		if (option.contains_ipv4) {
			if (containsIPv4(val)) {
				if (option.extract_ipv4) {
					out = append(out, extractIPv4(val)...)
				} else {
					out = append(out, val)
				}
			}
		}
		if (option.ipv4_alike) {
			if (containsIPv4alike(val)) {
				if (option.extract_ipv4) {
					if (option.as_ipv4) {
						out = append(out, transformIPv4alike(extractIPv4alike(val))...)
					} else {
						out = append(out, extractIPv4alike(val)...)
					}
				} else {
					out = append(out, val)
				}
			}
		}
		if (isIPv4(val)) {
			out = append(out, val)
		}
	// reverse match (non ipv4)
	} else {
		if (isIPv4(val)) {
			//continue
			return out
		}
		if (option.contains_ipv4 && containsIPv4(val)) {
			//continue
			return out
		}
		if (option.ipv4_alike && containsIPv4alike(val)) {
			//continue
			return out
		}
		out = append(out, val)
		//fmt.Fprintf(os.Stdout, "containsIPv4alike: '%s'\n", val)
	}
	return out
}

func isIPv4(val string) bool {
	pattern := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"

	re := regexp.MustCompile(pattern)
	return re.MatchString(val)
}
func containsIPv4(val string) bool {
	pattern := "(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])"

	re := regexp.MustCompile(pattern)
	return re.MatchString(val)
}
func containsIPv4alike(val string) bool {
	pattern := "(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])-){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])"

	re := regexp.MustCompile(pattern)
	return re.MatchString(val)
}

func extractIPv4(val string) []string {
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	submatchall := re.FindAllString(val, -1)
	out := []string{}
	for _, element := range submatchall {
		out = append(out, element)
	}
	return out
}
func extractIPv4alike(val string) []string {
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(-(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	submatchall := re.FindAllString(val, -1)
	out := []string{}
	for _, element := range submatchall {
		out = append(out, element)
	}
	return out
}

func transformIPv4alike(val []string) []string {
	out := []string{}
	for _, e := range val {
		out = append(out, strings.Join(strings.Split(e, "-"), "."))
	}
	return out
}

func getInput() *bufio.Scanner {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return bufio.NewScanner(os.Stdin)
	}
	return bufio.NewScanner(strings.NewReader(flag.Arg(0)))
}

func init() {
	flag.BoolVar(&option.contains_ipv4, "g", false, "")
	flag.BoolVar(&option.contains_ipv4, "contains-ipv4", false, "")

	flag.BoolVar(&option.ipv4_alike, "G", false, "")
	flag.BoolVar(&option.ipv4_alike, "ipv4-alike", false, "")

	flag.BoolVar(&option.extract_ipv4, "o", false, "")
	flag.BoolVar(&option.extract_ipv4, "extract-ipv4", false, "")

	flag.BoolVar(&option.invert_match, "v", false, "")
	flag.BoolVar(&option.invert_match, "invert-match", false, "")

	flag.BoolVar(&option.as_ipv4, "T", false, "")
	flag.BoolVar(&option.as_ipv4, "as-ipv4", false, "")

	flag.BoolVar(&option.debug, "d", false, "")
	flag.BoolVar(&option.debug, "debug", false, "")

	flag.Usage = func() {
		h := "Grep IP (GOLANG) - Extract / Exclude IPv4 addresses\n"
		h += "supports stdin/stdout for workflow integration\n\n"

		h += "Usage:\n"
		h += "  grepip (<string>) [options]\n\n"

		h += "Options:\n"
		h += "  -g, --contains-ipv4   contains IPv4, otherwise exact\n"
		h += "  -G, --ipv4-alike      contains IPv4 alike\n"
		h += "  -o, --extract         extract IPv4 match\n"
		h += "  -v, --invert-match    exclude IPv4 match\n"
		h += "  -T, --as-ipv4         transform IPv4 alike as IPv4\n"
		h += "  -d, --debug           enable debug output\n\n"

		h += "Examples:\n"
		h += "  grepip 'string 127.0.0.1#' [options]\n"
		h += "  cat data.txt | grepip [options]\n"
		h += "  cat data.txt | grepip [options]\n"

		fmt.Fprint(os.Stderr, h)
	}

}
