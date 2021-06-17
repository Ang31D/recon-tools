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

// curl -s https://unit42.paloaltonetworks.com/digital-quartermaster-scenario-demonstrated-in-attacks-against-the-mongolian-government/

// TODO
// # data input
// * as argument: -f <file>, -u <url>
//   if arg value is missing (ex -f <empty>) and stdin exists then stdin is arg type (file|url)
//   * item in stdin
//     - starts with 'http' = source.type: url
//     - file exist = source.type: file
//     - else = data
// in this way we know the source where we got the IOCs
//
// * ioc_source
//   type: file|url|data
//   path: filepath|url|<empty>
// data where source is unknown, could be from Stdin

/*
replace
\\r\\n # new line
\\t # tab

jq '[.[] | select(.type=="domain") | {type: .type, ioc: .ioc}]'

// ioc == domain
// ioc != defanged
// ioc does not contain (paloaltonetworks)
jq '[.[] | select(.type=="domain")]'
jq '[.[] | select(.defanged==false)]'

jq '[ .[] | select( .ioc | contains("paloaltonetworks") | not) ]'
jq '[ .[] | select( .ioc | contains("paloaltonetworks")) ]'

jq '[ .[] | select( .rootdomain | contains("paloaltonetworks")|not) ]'

jq '[.[] | {ioc: .ioc, data: .data}]'

jq '[ .[] | select(.ioc=="paloaltonetworks") ]'

cat web.content.2 | ./gioc | jq '[.[] | .ioc] | unique[]' | sed 's/"//g'


// IOC
cdaklle.housejjk.com
celeinkec.com
cocolco.com
dolimy.celeinkec.com
housejjk.com
jowwln.cocolco.com
ofhloe.com
pagbine.ofhloe.com
pplime.savecarrots.com
question.eboregi.com
question.erobegi.com
thbaw.ofhloe.com

ofhloe[.]com
celeinkec[.]com
eboregi[.]com
savecarrots[.]com
cocolco[.]com
housejjk[.]com
erobegi[.]com


// detected as defanged=true
thbaw.ofhloe.com
dolimy.celeinkec.com
question.eboregi.com
pplime.savecarrots.com
cocolco.com
ofhloe.com
housejjk.com
question.erobegi.com
celeinkec.com
pagbine.ofhloe.com
jowwln.cocolco.com
cdaklle.housejjk.com

// detected as default=false
api.w.org
app-guse4001.marketo.com
assets.adobedtm.com
bpo.gov.mn
browsehappy.com
energy.gov.mn
excite.co.jp
github.com
gomakethings.com
kasperskycontenthub.com
masm.gov.mn
mod.gov.mn
object.prototype.tostring.call # FP
prismacloud.io
s.w.org
schema.org
style.id # FP
t.target # FP
twitter.com
www.facebook.com
www.fireeye.com
www.google.com
www.linkedin.com
www.politik.mn
www.reddit.com
www.w3.org
xmlhttp.open # FP
yahoo.com
yoast.com
*/


type IOC_Type string

const (
	TypeIPv4	IOC_Type = "ip4"
	TypeIPv6		 = "ip6"
	TypeDomain		 = "domain"
	TypeURL			 = "url"
	TypeEmail		 = "email"
	TypeMD5			 = "md5"
	TypeSHA1		 = "sha1"
	TypeSHA256		 = "sha256"
)

type IOC struct {
	Value string	`json:"ioc"`
	Type  IOC_Type	`json:"type"`
	Defanged bool	`json:"defanged"`
	RootDomain string `json:"rootdomain"`
	Data   string	`json:"data"`
}

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
			if (seen[ioc.Value]) {
				continue
			}

			ioc_list = append(ioc_list, ioc)
			seen[ioc.Value] = true
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
		out = append(out, IOC{Value: domain, Type: TypeDomain, Data: data, Defanged: true, RootDomain: getRootDomain(domain)})
	}
	return out
}

func extractDomains(data string) []IOC {
	out := []IOC{}
	result := xioc.ExtractDomains(data)
	for _,domain := range result {
		has_defang := hasDotDefang(data)
		domain = strings.ToLower(domain)
		// hasDotDefang check the whole data,  triggers FP (true) on '@'
		// check if non '@<domain>' is found in data
		if (strings.Contains(strings.ToLower(data), domain)) {//
			has_defang = false
		}
		out = append(out, IOC{Value: domain, Type: TypeDomain, Data: data, Defanged: has_defang, RootDomain: getRootDomain(domain)})
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
