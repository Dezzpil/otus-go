# ocharb

Often characters beyond ( the ranges )

## default config:
```
0000-007f // latin range
0410-044f // subrange from cyrillic range
0401 // ё
0451 // Ё
```

## uses:
```./ocharb < file.txt``` 

## example:
```
$ echo "Ludwig van Beethoven K 125. Výroćí Jeno Úmrtí : Symfonie" | ./ocharb
unicode;character;count
"U+00FD";"ý";1
"U+0107";"ć";1
"U+00ED";"í";2
"U+00DA";"Ú";1

```

## todo:
* read text from stdin
* output line of characters per line
* arguments parsing
* output for --as-csv
* output for --as-json
* write info to stderr
* parse config (default config above)
* init with user config, given with --config