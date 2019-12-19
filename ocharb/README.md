#ocharb
Often characters beyond ( the ranges )

## default config:
```
0000-007f // latin range
0410-044f // subrange from cyrillic range
0401 // ё
0451 // Ё
```

##todo:
* read text from stdin
* output line of characters per line
* arguments parsing
* output for --as-csv
* output for --as-json
* write info to stderr
* parse config (default config above)
* init with user config, given with --config

## uses:
```ocarb < file.txt``` 
