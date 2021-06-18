# GIOC

## TODO
```
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
```

## Query with JQ
```
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
cat web.content | ../gioc | jq '[.[] | select(.defanged==true)][0]'
cat web.content | ../gioc | jq '[.[] | select(.defanged==true) | .ioc]|unique[]' | sed 's/"//g'
cat web.content | ../gioc | jq '[.[] | .ioc]|unique[]' | sed 's/"//g' | revasset | sort -V | revasset

cat data/web.content | ./gioc | jq '[.[] | select(.item.type=="domain")]' | jq '[.[] | select(.item.verified==true)|.item]'
```
