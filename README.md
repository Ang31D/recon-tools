# recon-tools
Tools for Recon

All tools has a help section by using the '-h' argument.

## example usage

### crt.sh - A utility for quickly searching presorted DNS names.
Usage: 'crtsh (<domain> | -f <file> -h, --help)'
* fetch google subdomains containing 'api'

crtsh %25api%25.google.com

### fetch subdomains
dnsgrep google.com

### exclude ip addresses from a file
grepip file -v

### output alive (sub-)domains listed in the file
cat file | hostprobe

### Reverse of each sub-name (ex. com.google)
echo google.com | rdns
