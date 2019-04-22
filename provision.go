package provision

import (
	"bytes"
	"text/template"
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

// Create server config and return bytes buffer to be written to file.
func CreateServer(name string, location []string) bytes.Buffer {
	var result bytes.Buffer
	tpl := template.Must(template.ParseFiles("templates/server.tpl"))

	server := Server{
		ServerName:  name,
		LogLocation: "/var/log/" + name + "-nginx.log",
		Locations:   location,
	}
	err := tpl.Execute(&result, server)
	if err != nil {
		panic(err)
	}

	return result

}

// Create location string.
func CreateLocation(name string, proxy bool, proxyAddr string, isWebsocket bool) string {
	var loc bytes.Buffer

	tplFile, err := template.ParseFiles("templates/location.tpl")
	if err != nil {
		panic(err)
	}

	if proxy {
		tplFile, err = template.ParseFiles("templates/location_proxy.tpl")
		if err != nil {
			panic(err)
		}
		location := Proxy{
			Location:     name,
			ProxyAddress: proxyAddr,
			IsWebsocket:  isWebsocket,
		}
		tpl := template.Must(tplFile, nil)
		err = tpl.Execute(&loc, location)
		if err != nil {
			panic(err)
		}
	} else {
		location := Location{
			Location: name,
			Root:     "/var/www/html" + name,
		}
		tpl := template.Must(tplFile, nil)
		err = tpl.Execute(&loc, location)
		if err != nil {
			panic(err)
		}
	}
	return loc.String()

}
