## bandict

Command-line [urbanDictionary](https://www.urbandictionary.com/) Search

### Change Preferences

Three constants have been defined in `main.go`. Change them to change preferences.

```go
NO_OF_RESULTS  = 1
WRAP_WIDTH     = 80
DISPLAY_FOOTER = true
```

**NO_OF_RESULTS**: Number of definitions to be displayed by default (optionally can be specified with the `-n` flag).

**WRAP_WIDTH**: Word-wrap width for the output text.

**DISPLAY_FOOTER**: Display footer containing number of results fetched and displayed.


### Flags

`-n`: Number of results to be displayed. If not specified, **NO_OF_RESULTS** is used.

`-w`: Search string. If not specified, requested in the next line.

`-s`: Display sound files for the search string (not the definitions).


### Examples

Display three results. Search string requested in the next line.

```bash
$ bandict -n 3
```

Display three results for the string "thug life".

```bash
$ bandict -w 'thug life' -n 3
```

Display sound files (pronunciations) for the search string. One on each line.

```bash
$ bandict -w 'lmao' -s
```

Use the command to download the first file:

```bash
$ wget $(bandict -w 'lmao' -s | sed -n 1p)
```

### License

See LICENSE file.

