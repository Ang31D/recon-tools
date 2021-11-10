#!/bin/bash

if [[ -p /dev/stdin ]]; then
  data=$(cat "/dev/stdin")
elif [[ $# -gt 0 ]]; then
  data=$(echo $1)
fi

echo "$data" | sed 's/^hxxp/http/g' | sed 's/\[\.\]/\./g' |sed 's/\[:\/\/\]/:\/\//g' | sed 's/\[:\]\/\//:\/\//g'
