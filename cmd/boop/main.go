package main

import "fmt"

const (
	banner = `
	 ____   __    __  ____ 
	(  _ \ /  \  /  \(  _ \
	 ) _ ((  O )(  O )) __/
	(____/ \__/  \__/(__) 
	`
	usage = `
	Usage:
	-url
		Target HTTP server URL (required)
	-n 
		Number of requests to make
	-c
		Concurrency level
	`
)

func showBanner() string { return banner[1:] }
func showUsage() string { return usage[1:] }

func main() {
	fmt.Println(showBanner())
	fmt.Println(showUsage())
}
