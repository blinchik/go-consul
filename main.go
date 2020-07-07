package main

import (
	"flag"
	"log"
	"os"

	discover "github.com/blinchik/go-utils/discover"
	acl "gitlab.com/blinchik-go/go-consul/acl"
)

var consulAddress string
var consulPort = "8500"

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	consulAddress = discover.GetIP("Name", "blinchik_consul")
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bootstrapACL := flag.Bool("bootstrap", false, "bootstrap")

	flag.Parse()

	if *bootstrapACL {

		if os.Args[2] != "" {

			consulAddress = os.Args[2]
			consulPort = "80"
		}

		acl.BootstrapACL(consulAddress, consulPort)

	}

}
