package provision

import (
	"bytes"
	"html/template"
	"os"
)

type Server struct {
	ServerName  string
	LogLocation string
	Locations   []string
}

type Location struct {
	Location string
	Root     string
}

type Proxy struct {
	Location     string
	ProxyAddress string
	IsWebsocket  bool
}

func ProvisionServer(names ...string) {
	tpl := template.Must(template.ParseFiles("templates/server.tpl"))
	locTpl := template.Must(template.ParseFiles("templates/location.tpl"))
	for i := 0; i < len(names); i++ {
		var loc bytes.Buffer
		resultFile, err := os.Create("nginx-" + names[i])
		if err != nil {
			panic(err)
		}
		rootLocation := Location{
			Location: "/",
			Root: "/var/www/html",
		}
		err = locTpl.Execute(&loc, rootLocation)
		server := Server{
			ServerName: names[i],
			LogLocation: "/var/log/" + names[i] + "-nginx.log",
			Locations: []string{
				loc.String(),
			},
		}
		err = tpl.Execute(resultFile, server)
		if err != nil {
			panic(err)
		}
	}

}

func ProvisionProxy(name string, proxy Proxy) {

}