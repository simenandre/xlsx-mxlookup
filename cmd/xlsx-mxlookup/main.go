package main

import (
	"flag"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/tealeg/xlsx"
)

var suppliers map[string]*regexp.Regexp

type channelInfo struct {
	RowInt int
	Host   string
}

func init() {
	suppliers = make(map[string]*regexp.Regexp)
	suppliers["google"] = regexp.MustCompile(`(?m)google\.com`)
	suppliers["office365"] = regexp.MustCompile(`(?m)onmicrosoft\.com`)
	suppliers["office365"] = regexp.MustCompile(`(?m)outlook\.com`)
}

func main() {
	file := flag.String("input", "", "file path. eg. ./fixtures/domain-test.xlsx")
	o := flag.String("output", "./output.xlsx", "eg. ./output.xlsx (defaults to output.xlsx)")
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
	rows := len(sh.Rows)
	bar := pb.StartNew(rows)
	var wg sync.WaitGroup
	ch := make(chan channelInfo, rows)

	for i, r := range sh.Rows {
		e := r.Cells[*c]
		wg.Add(1)
		lookupMx(e.String(), i, ch, &wg, bar)
	}

	wg.Wait()
	close(ch)

	for msg := range ch {
		if msg.Host != "" {
			sh.Rows[msg.RowInt].AddCell().SetString(msg.Host)
		}
	}

	wb.Save(*o)
}


func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}
 
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func lookupMx(domain string, rowInt int, ch chan channelInfo, wg *sync.WaitGroup, bar *pb.ProgressBar) {
	defer wg.Done()

	var done = func(k string) {
		chi := channelInfo{
			RowInt: rowInt,
			Host:   k,
		}

		bar.Increment()

		ch <- chi

	}

	retry(3, time.Second, func() error {
		var r string
		mxrecords, err := net.LookupMX(domain)
		if err != nil {
			return stop{err}
		}

		for _, mx := range mxrecords {
			r = r + mx.Host + ","
		}

		for key, t := range suppliers {
			if t.MatchString(strings.ToLower(r)) {
				done(key)
				return nil
			}
		}

		done(r)
		return nil
	})
}
