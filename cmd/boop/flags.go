package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	url         string
	num, clevel int
}

const usageText = `
Usage:
	boop [options] url
Options:
`

func (f *flags) parse() error {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageText[1:]) 
        flag.PrintDefaults()
	}
	
	flag.Var(toNumber(&f.num), "n", "Number of requests to make")                 
	flag.Var(toNumber(&f.clevel), "c", "Concurrency level")                
	flag.Parse()

	f.url = flag.Arg(0)	
	if err := f.validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		return err
	}
	return nil
}

func (f *flags) validate() error {
	if f.clevel > f.num {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.clevel, f.num)
	}
	if err := validateURL(f.url); err != nil {
		return fmt.Errorf("url: %w", err)
	}
	return nil
}

func validateURL(s string) error {
    u, err := url.Parse(s)   
    switch {
    case strings.TrimSpace(s) == "":
        err = errors.New("required")
    case err != nil:         
        err = errors.New("parse error")         
    case u.Scheme != "http": 
        err = errors.New("only supported scheme is http") 
    case u.Host == "":       
        err = errors.New("missing host")        
    }
    return err 
}

type number int

func toNumber(p *int) *number {
	return (*number)(p)
}

func (n *number) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	switch {
	case err != nil:
        err = errors.New("parse error")
    case v <= 0:
        err = errors.New("should be positive")
    }
    *n = number(v) 
    return err
}
 
func (n *number) String() string {
    return strconv.Itoa(int(*n))      
}