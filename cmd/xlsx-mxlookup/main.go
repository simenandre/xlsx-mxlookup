package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/tealeg/xlsx"
)

var suppliers map[string]*regexp.Regexp

func init() {
	suppliers = make(map[string]*regexp.Regexp)
	suppliers["google"] = regexp.MustCompile(`(?m)google\.com`)
	suppliers["office365"] = regexp.MustCompile(`(?m)onmicrosoft\.com`)
	suppliers["office365"] = regexp.MustCompile(`(?m)outlook\.com`)
}

func main() {
	file := flag.String("input", "", "file path. eg. ./fixtures/domain-test.xlsx")
	o := flag.String("output", "./test.xlsx", "eg. ./output.xlsx (defaults to output.xlsx)")
	c := flag.Int("col", 6, "describe what column to read domain from")
	flag.Parse()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *o == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	wb, err := xlsx.OpenFile(*file)
	if err != nil {
		panic(err)
	}

	sh := wb.Sheets[0]
	for _, r := range sh.Rows {
		e := r.Cells[*c]
		host, err := lookupMx(e.String())
		if err == nil {
			r.AddCell().SetString(host)
		}
	}

	fmt.Println(*o)

	wb.Save(*o)
}

func lookupMx(domain string) (string, error) {
	var r string
	mxrecords, err := net.LookupMX(domain)
	if err != nil {
		return "", err
	}

	for _, mx := range mxrecords {
		r = r + mx.Host + ","
		// fmt.Println(mx.Host, mx.Pref)
	}

	for key, t := range suppliers {
		if t.MatchString(strings.ToLower(r)) {
			return key, nil
		}
	}

	return r, nil
}
