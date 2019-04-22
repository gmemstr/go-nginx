package main

import (
	"github.com/gmemstr/nginx/provision"
)

func main() {
	var locations []string
	location := provision.CreateLocation("/test", false, "", false)
	location2 := provision.CreateLocation("/test2", true, "http://127.0.0.1:4343", false)
	location3 := provision.CreateLocation("/test3", true, "http://127.0.0.1:4343", true)

	locations = append(locations, location, location2, location3)
	provision.CreateServer("example.com", locations)
}
