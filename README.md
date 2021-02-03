<img width="400" src=".github/assets/header.svg">

![GitHub Release Date](https://img.shields.io/github/release-date/cobraz/xlsx-mxlookup)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/cobraz/xlsx-mxlookup.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/cobraz/xlsx-mxlookup/alerts/)
[![Maintainability](https://api.codeclimate.com/v1/badges/9b7914dc64de7eb7e76f/maintainability)](https://codeclimate.com/github/cobraz/xlsx-mxlookup/maintainability)

This utility is written in Go and features a simple way to loop through a list of domains and return their MX records. Usecase is typically for sales, to easily get what MX record domains have.

<img width="500" src=".github/assets/recording.svg">

## Installation

Using Mac?

```shell
brew install cobraz/tools/xlsx-mxlookup
```

**Notes**: The library is not tested on Linux or Windows. There are [executables available](https://github.com/cobraz/xlsx-mxlookup/releases/latest).

## Help

```shell
# Example command
▶ xlsx-mxlookup --input ./domains.xlsx --col 7
```

```shell
▶ xlsx-mxlookup --help
Usage of xlsx-mxlookup:
  -col int
    	describe what column to read domain from (default 6)
  -input string
    	file path. eg. ./fixtures/domain-test.xlsx
  -output string
    	eg. ./output.xlsx (defaults to output.xlsx) (default "./output.xlsx")
```

`--input` is required and should point to your file. `col` is not required, but if you don't have your domain in column number 7, you'll have problems. You can also define where to store your XLSX output, with
the `--output` command.

## Contribute

Please, oh pretty please do contribute! If you feel this helps you out, but you want to increase the quality of this software, please submit pull requests. Look at our issues page for more information – as previously stated, it's pretty bare-bone. Making it faster, better is something everyone wants.
