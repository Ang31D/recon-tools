package main

import (
	"fmt"
	"bufio"
	"flag"
	"os"
	"strings"
	"encoding/json"
	"regexp"

	"github.com/assafmo/xioc/xioc"
)

type IOC_ItemType string

const (
	IOC_TypeIPv4	IOC_ItemType = "ip4"
	IOC_TypeIPv6	 = "ip6"
	IOC_TypeDomain	 = "domain"
	IOC_TypeURL		 = "url"
	IOC_TypeEmail	 = "email"
	IOC_TypeMD5		 = "md5"
	IOC_TypeSHA1	 = "sha1"
	IOC_TypeSHA256	 = "sha256"
)

type IOC_SourceType string

const (
	Source_TypeUnknown	IOC_SourceType = "unknown"
	Source_TypeFile	= "file"
	Source_TypeURL	= "url"
)


type IOC_Tag struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type IOC_Details struct {
	Source string `json:"source"`
	SourceType IOC_SourceType `json:"source_type"`
	Data string `json:"data"`
	HasDefang bool `json:"has_defang"`
}

type IOC_Item struct {
	Value string `json:"value"`
	Type  IOC_ItemType	`json:"type"`
	Verified  bool	`json:"verified"` // Verified_IOC
	Tags []IOC_Tag	`json:"tags"`
}
/*
type IOC_Item_TypeDomain struct {
	FQDN string `json:"value"`
	RootDomain string `json:"rootdomain"`
}
*/
type IOC struct {
	Item IOC_Item `json:"item"`
	Details IOC_Details `json:"details"`
}
/*
type IOC struct {
	Value string	`json:"item"`
	Type  IOC_Type	`json:"type"`
	Is_IOC  bool	`json:"is_ioc"`
	Data_Defanged bool	`json:"data_defanged"`
	RootDomain string `json:"rootdomain"`
	Source_Data   string	`json:"source_data"`
	Details IOC_Details 	`json:"debug"`
}
*/

func main() {
	sc := getInput()

	var ioc_list []IOC

	seen := make(map[string]bool)
	//seen_defang := make(map[string]bool)

	for sc.Scan() {
		data := formatEscapedData(sc.Text())

		//result := extractDefangDomains(data)
		result := extractDomains(data)
		if (len(result) == 0) {
			continue
		}

		for _, ioc := range result {
			//if (seen[ioc.Value]) {
			if (seen[ioc.Item.Value]) {
				continue
			}

			ioc_list = append(ioc_list, ioc)
			//seen[ioc.Value] = true
			seen[ioc.Item.Value] = true
		}
	}

	iocJsonBlob := iocListAsJsonBlob(ioc_list)
	fmt.Fprintf(os.Stdout, "%s\n", iocJsonBlob)
}

func formatEscapedData(data string) string {
	out := strings.ReplaceAll(data, "\\r\\n", "\n")
	out = strings.ReplaceAll(out, "\\n", "\n")
	out = strings.ReplaceAll(out, "\\t", "\t")
	out = strings.ReplaceAll(out, "\\/", "/")
	return out
}

func iocListAsJsonBlob(ioc_list []IOC) string {
	var ioc_json []string
	for _, ioc := range ioc_list {
		jsonIOC, _ := json.Marshal(ioc)
		ioc_json = append(ioc_json, string(jsonIOC))
	}
	return fmt.Sprintf("[%s]", strings.Join(ioc_json, ","))
}

//var dot = `(\.|\p{Z}dot\p{Z}|\p{Z}?(\(dot\)|\[dot\]|\(\.\)|\[\.\]|\{\.\})\p{Z}?)`
var dot = `(\p{Z}dot\p{Z}|\p{Z}?(\(dot\)|\[dot\]|\(\.\)|\[\.\]|\{\.\})\p{Z}?)`
var dotRegex = regexp.MustCompile(`(?i)` + dot)
func hasDotDefang(data string) bool {
	out := false
	if (dotRegex.MatchString(data)) {
		out = true
	}
	return out
}

func getRootDomain(domain string) string {
	return reverseDomain(strings.Join(strings.Split(reverseDomain(domain), ".")[:2], "."))
}
func reverseDomain(domain string) string {
        parts := strings.Split(domain, ".")
        for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
                parts[i], parts[j] = parts[j], parts[i]
        }
        return strings.Join(parts, ".")
}

/*
var domainRegex = regexp.MustCompile(`(?i)([\p{L}\p{N}][\p{L}\p{N}\-]*` + dot + `)+\p{L}{2,}`)
//var domainRegex = regexp.MustCompile(`([\p{L}\p{N}][\p{L}\p{N}\-]*` + dot + `)+\p{L}{2,}`)
// ISSUE: only extracts defanged rootdomain, no subdomain due to removed regular '.' from 'dot' regex pattern
func extractDefangDomains(data string) []IOC {
	//out := []IOC{}
	out := []IOC{}

	result := domainRegex.FindAllString(data, -1)
	if len(result) == 0 {
		return out
	}
	//for _, domain := range result {
	for _, domain := range result {
		domain = strings.ToLower(domain)
		//fmt.Fprintf(os.Stdout, "%d: %s\n", i, strings.ToLower(domain))
		//out = append(out, IOC{Value: domain, Type: IOC_TypeDomain, Source_Data: data, Defanged: true, RootDomain: getRootDomain(domain)})
		out = append(out, IOC{Value: domain, Type: IOC_TypeDomain, Source_Data: data,  Data_Defanged: true, RootDomain: getRootDomain(domain), Is_IOC: true})
	}
	return out
}
*/

func extractDomains(data string) []IOC {
	out := []IOC{}
	result := xioc.ExtractDomains(data)
	for _,domain := range result {
		has_defang := hasDotDefang(data)
		is_ioc := false
		domain = strings.ToLower(domain)
		// hasDotDefang check the whole data,  triggers FP (true) on '@'
		// check if non '@<domain>' is found in data
		if (!strings.Contains(strings.ToLower(data), domain)) {
			is_ioc = true
		}
		var ioc_tags []IOC_Tag
		ioc_item := IOC_Item{Value: domain, Type: IOC_TypeDomain, Verified: is_ioc, Tags: ioc_tags}
		ioc_details := IOC_Details{Source: "", SourceType: Source_TypeUnknown, Data: data, HasDefang: has_defang}
		ioc := IOC{Item: ioc_item, Details: ioc_details}
		out = append(out, ioc)
		
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
