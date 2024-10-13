package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/theredwiking/cacheproxy/pkg/origin"
	"github.com/theredwiking/cacheproxy/pkg/server"
)

func main() {
	port := flag.Int("port", 3000, "define what port to run proxy on")
	url := flag.String("origin", "", "what url to send requests to")
	flag.Parse()

	if *url == "" {
		fmt.Println("Missing url to send requests to")
		os.Exit(0)
	}

	origin := origin.NewOrigin(*url)
	server.Serve(*port, origin)
}
