package main

import (
	"flag"
	"log"
	"os"

	acl "github.com/blinchik/go-consul/acl"
)

var consulAddress string
var consulPort = "8500"
var consulRootPath string

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bootstrapACL := flag.Bool("bootstrap", false, "bootstrap")

	flag.Parse()

	if *bootstrapACL {

		if os.Args[2] != "" {

			consulAddress = os.Args[2]
			consulPort = os.Args[3]
			consulRootPath = os.Args[4]
		}

		acl.BootstrapACL(consulAddress,consulRootPath, consulPort)

	}

}
