# recon-tools
Tools for Recon

All tools has a help section by using the '-h' argument.

## example usage

### crtsh - A utility for quickly searching presorted DNS names.
```Usage: 'crtsh (<domain> | -f <file> -h, --help)'```
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
```Usage: 'hostprobe (<dns> | -f <file>) [options] -h, --help'```
* output alive (sub-)domains listed in the file

```hostprobe www.sites.google.com```

```cat <dns_file> | hostprobe```

### rdns - Reverse of each sub-name (ex. com.google).
```Usage: 'rdns (<file> | -r <dns>) -h, --help'```
* Reverse each sub-name (ex. com.google)

```rdns google.com```

```cat <dns_list> | rdns```

### rhead2json - A utility to convert HTTP response header to json format.
```Usage: 'rhead2json (<url>) -h, --help'```
* Fetch headers from https://www.google.com as json-format

```rhead2json https://www.google.com```

```curl -s -IXGET https://www.google.com | tee response.raw | rhead2json | tee response.headers.json```
