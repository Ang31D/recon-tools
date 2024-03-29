# recon-tools
Tools for Recon

All tools has a help section by using the '-h' argument.

All tools should also support stdin/out for workflow integration.

## example usage

### crtsh - crtsh - utility to search for domain(s) through crt.sh (presorted DNS names)
```
crtsh - utility to search for domain(s) through crt.sh (presorted DNS names)
stdin/out support for workflow integration
Usage: 'crtsh (<domain>|<company>) [options]'
  query:
    default query sub-domains if omitting: --tld, --cn, --dns or --org
  --tld                   matching on any tld ('<domain-name>.*' instead of '*.<domain>')
                          hint: useful when finding root or 3rd-party domains
  --cn                    match on common name (Subject commonName)
  --dns                   match on dns (Subject Alternative Name)
  --org                   match on company (Subject organizationName)
  --identity              match on identity (matching on 'Full Text Search'

  output:
    default output domains if omitting: -p, -r or -s
    -w, --strip-wildcard    strip wildcard ('*.') from dynamic dns name results
    -a, --append            append ',<name_value>' to <common_name>,
                            with '--org' <name_value> will be the <company>,
                            with '--cn' <common_name> does not exist, will only output <name_value>
    -A, --rev-append        reverse append ',<common_name>' to <name_value>,
    -p, --pretty-json       output as pretty-json
    -r, --raw-json          output as raw-json
    -s, --json-stream       output one json blob per line
                            hint: useful when looking at the result format or stdin/out workflow

  input:
    -J, --json-input        stdin is raw json (from previous '-r' output)
```
* fetch google subdomains containing 'api'

```crtsh %25api%25.google.com```

```cat <dns_list> | crtsh```

### dnsgrep - A utility for quickly searching presorted DNS names.
```Usage: 'dnsgrep (<dns> | -f <file>) [options] -h, --help'```
* fetch google subdomains

```dnsgrep google.com```

```cat <dns_list> | dnsgrep```

### grepip - Pull / Exclude IPv4 addresses
```Usage: 'grepip (<file>) [options] -h, --help'```
* exclude ip addresses from a file

```grepip <result_from_dnsgrep> -v```

```dnsgrep google.com -L -F | grepip -v```

### hostprobe - A utility for quickly see if a dns/host exists.
```
stdin/out support for workflow integration
Usage: 'hostprobe (<host>) [options]'
  <host>                   the host (domain|ipv4|ipv6) to lookup
  -s, --dns-server         dns server to query (optional)
  -t, --type               specifies the query type (optional), ex. CNAME, NS, SOA, TXT, DNSKEY, AXFR, etc.
  -n, --not-found          only output hosts not found
  -A, --allow-ip           allow probing ip-address
  -r, --raw                outputs the raw results (supports -N, -v)
  -N, --exclude-not-found  exclude hosts not found, requires -r
  -F, --filter             TO IMPLEMENT: filter raw output (requires -r)
  -x, --resolve            resolve host (supports -f)
  -f --format              outputs "<host> <resolved host,...>" (requires -x)
  -E, --extract-metadata   extract dns/ip (supports -f)
  -v, --verbose            verbose output
  -vv                      extra verbose output
  -mD, --meta-domain       append ,<domain> for ip
  -m <d,i,m|all>, --meta <d,i,m|all> append ,Alias for domain ,Domain for ip, Ip for domain, Mail or All
    meta options
    d, domain-alias        return string containing 'is an alias for'
```
* output alive (sub-)domains listed in the file

```hostprobe www.sites.google.com```

```cat <dns_file> | hostprobe```

### revasset - Reverse of each sub-name (ex. com.google).
```Usage: 'revasset (<dns> | <ip>) -h, --help'```
* Reverse each sub-name (ex. com.google)

```revasset google.com```

```cat <dns_list> | revasset```

### rhead2json - A utility to convert HTTP response header to json format.
```Usage: 'rhead2json (<url>) -h, --help'```
* Fetch headers from https://www.google.com as json-format

```rhead2json https://www.google.com```

```curl -s -IXGET https://www.google.com | tee response.raw | rhead2json | tee response.headers.json | jq```

### urigrep - A utility for quickly return uri/url links.
```Usage: 'urigrep (<url>) -h, --help'```
* Fetch uri/urls from https://www.google.com

```urigrep https://www.google.com```

```curl -s https://www.google.com | urigrep```

### gurl - Grep Url - Parse url.
```Grep Url - Parse url
Usage: 'gurl [<file>] [-b|-p|-d]'
  -p, --parse-uri-path        (default) output uri '</path>' (exclude 'http(s)://<host-dns>')
  -b, --parse-base-url        output root-url 'http(s)://<host-dns>' (exclude '</path>*')
  -d, --parse-domain          output unique host dns (exclude 'http(s)://' and '</path>*')
  -s, --show-stats            show number of occurrence and sequence
  -r, --reverse-stats         sort stats number of occurrence in reverse
  -f, --filter [<base-url>]   filter on base-url (see 'filter options')
    filter options
    -n, --stats-number <number>   filter on (base-url) number from the '-s -b' output
```

### hidb - Manage hash-based index db.
```Hash Index - Manage hash-based index db.
Usage: 'hidb (<resource>) [-o|-I|-N|-L] [options]'
  -h, --help                      display this help and exit
  -o, --output <folder>           Write index to output folder (default: ./)
  -w, --write                     Write resource file (md5-files, combine with -o).
  -I, --indexed <folder>          Filter by indexed resources (md5 file)
  -F, --index-file <yes/no>       Check if resource (md5) file exists (yes) or not (no)
  -N, --no-index <folder>         Filter resources not in index (md5 file)
  --hash                          Include hash in output (supports: -I, -N, -L)
  --hash-only                     Only output hash (supports: -I, -N)
  -L, --lookup <folder>           Display resource (from index file) for index (md5 hash)
  -n, --position                  Include position (row number) in index file
  --cleanup                       Cleanup orphan indexes (md5-file not in index)
  -D, --dry-run                   Dry, do not store in files.
  ```
  
  ### cidb - Curl indexed url - from hash-based index db.
```Curl indexed url - from hash-based index db.
Usage: 'cidb (<md5-hash>:<url>) ([-i|-D|-v]) [options]'
  -h, --help                        display this help and exit
  -o, --output <folder>             Write index to folder (default: ./)
  -i, --include                     Include protocol response headers in the output
  -D, --dump-header <folder>        Write the received headers to <folder> (default: ./response_header/)
  -d, --dump-header-suffix <file>   Write the received headers to file (<hash>.<suffix>) in (-o) folder)
  -v, --verbose                     Make the operation more talkative
  ```
  
  ### gidb - Grep hash-based index db.
```Grep Hash Index - Grep hash-based index db.
Usage: 'gidb (<file>) -hS|-hH (<folder>) [options]'
  -h, --help                    display this help and exit
  -hS, --http-status <folder>   Grep HTTP status in folder (default: .)
  -hH, --http-header <folder>   Grep HTTP header in folder (default: .)
  ```
  
