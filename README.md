# recon-tools
Tools for Recon

All tools has a help section by using the '-h' argument.

## example usage

### crtsh - A utility for quickly searching presorted DNS names.
```Usage: 'crtsh (<domain> | -f <file> -h, --help)'```
* fetch google subdomains containing 'api'

```crtsh %25api%25.google.com```

```cat <dns_list> |crtsh```

### dnsgrep - A utility for quickly searching presorted DNS names.
```Usage: 'dnsgrep (<dns> | -f <file>) [options] -h, --help'```
* fetch google subdomains

```dnsgrep google.com```

### grepip - Pull / Exclude IPv4 addresses
```Usage: 'grepip (<file>) [options] -h, --help'```
* exclude ip addresses from a file

```grepip <result_from_dnsgrep> -v```

### hostprobe - A utility for quickly see if a dns/host exists.
```Usage: 'hostprobe (<dns> | -f <file>) [options] -h, --help'```
* output alive (sub-)domains listed in the file

```cat <dns_file> | hostprobe```

### rdns - Reverse of each sub-name (ex. com.google).
```Usage: 'rdns (<file> | -r <dns>) -h, --help'```
* Reverse each sub-name (ex. com.google)

```echo google.com | rdns```
