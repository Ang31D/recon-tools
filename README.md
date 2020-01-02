# recon-tools
Tools for Recon

All tools has a help section by using the '-h' argument.

## example usage

### crtsh - A utility for quickly searching presorted DNS names.
```Usage: 'crtsh (<domain> | -f <file> -h, --help)'```
* fetch google subdomains containing 'api'

```crtsh %25api%25.google.com```

### dnsgrep - A utility for quickly searching presorted DNS names.
```Usage: 'dnsgrep (<dns> | -f <file>) [options] -h, --help'```
* fetch subdomains

```dnsgrep google.com```

### Grep IP - Pull / Exclude IPv4 addresses
```Usage: 'grepip (<file>) [options] -h, --help'```
* exclude ip addresses from a file

```grepip <result_from_dnsgrep> -v```

### output alive (sub-)domains listed in the file
cat file | hostprobe

### Reverse of each sub-name (ex. com.google)
echo google.com | rdns
