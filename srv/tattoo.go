package main

import (
	"flag"
	"github.com/xyproto/tattoo"
)

func main() {
	var useFCGI = flag.Bool("fcgi", false, "Use FastCGI")
	flag.Parse()
	tattoo.Serve(*useFCGI)
}
