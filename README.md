# XLSX MXLookup

This utility is written in Go and features a simple way to loop through a list of domains and return their MX records. Usecase is typically for sales, to easily get what MX record domains have.

## Installation

Using Mac?

```shell
brew install cobraz/tools/xlsx-mxlookup
```

**Notes**: The library is not tested on Linux or Windows. There are [executables available](https://github.com/cobraz/xlsx-mxlookup/releases/latest).

## Help

```shell
# Example command
xlsx-mxlookup --input ./domains.xlsx --col 7
```

`--input` is required and should point to your file. `col` is not required, but if you don't have your domain in column number 7, you'll have problems. You can also define where to store your XLSX output, with
the `--output` command.

## Contribute

Please, oh pretty please do contribute! If you feel this helps you out, but you want to increase the quality of this software, please submit pull requests. Look at our issues page for more information – as previously stated, it's pretty bare-bone. Making it faster, better is something everyone wants.
