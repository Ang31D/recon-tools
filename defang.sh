#!/bin/bash

if [[ -p /dev/stdin ]]; then
  data=$(cat "/dev/stdin")
elif [[ $# -eq 1 ]]; then
  data=$(echo $1)
else
  exit
fi

echo "$data" | sed 's/^http/hxxp/g' | sed 's/\./\[\.\]/g' | sed 's/:\/\//\[:\/\/\]/g'
