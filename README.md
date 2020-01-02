# recon-tools
Tools for Recon

---- example ----
# fetch google subdomains containing 'api'
crtsh %25api%25.google.com

# fetch subdomains
dnsgrep google.com

# exclude ip addresses from a file
grepip <file> -v

# output alive (sub-)domains listed in the file
cat <file> | hostprobe

# Reverse of each sub-name (ex. com.google)
echo google.com | rdns
