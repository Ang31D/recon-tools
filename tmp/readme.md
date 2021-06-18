
```
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
```
