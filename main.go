package main

import (
	"flag"
	"log"
	"os"

	secret "github.com/blinchik/go-aws/lib/secrets"

	acl "github.com/blinchik/go-consul/acl"
)

var consulAddress string
var consulPort = "8500"
var consulRootPath string

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bootstrapACL := flag.Bool("bootstrap", false, "bootstrap")
	updateACL := flag.Bool("updateACL", false, "updateACL")

	flag.Parse()

	if *bootstrapACL {

		if os.Args[2] != "" {

			consulAddress = os.Args[2]
			consulPort = os.Args[3]
			consulRootPath = os.Args[4]
		}

		acl.BootstrapACL(consulAddress, consulRootPath, consulPort)

	}

	if *updateACL {

		if os.Args[2] == "global" {
			consulAddress = os.Args[3]
			consulPort = os.Args[4]
			consulRootPath = os.Args[5]

			secretValue := secret.GetSecret("global-management", "AWSCURRENT")

			acl.UpdateACLToken(consulAddress, consulRootPath, consulPort, secretValue, secretValue)

		} else {

			consulAddress = os.Args[2]
			consulPort = os.Args[3]
			consulRootPath = os.Args[4]

			secretValue := secret.GetSecret("global-management", "AWSCURRENT")

			acl.UpdateACLToken(consulAddress, consulRootPath, consulPort, os.Args[5], secretValue)

		}

	}

}
