package main

import (
	"github.com/gmemstr/nginx/provision"
	"os"
)

func main() {
	provision.ProvisionServer(os.Args[1])
}
