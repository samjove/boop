package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	banner = `
	 ____   __    __  ____ 
	(  _ \ /  \  /  \(  _ \
	 ) _ ((  O )(  O )) __/
	(____/ \__/  \__/(__) 
	`
)

func showBanner() string { return banner[1:] }

func main() {
	f := &flags{
		 num: 100,
		 clevel: runtime.NumCPU(),
	}

	if err := f.parse(); err != nil {
		os.Exit(1)
	}
	fmt.Println(showBanner())
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n", f.num, f.url, f.clevel)
	
}
