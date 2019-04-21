package main

import (
	"bytes"
	"github.com/gmemstr/nginx/provision"
	"io/ioutil"
	"os"
	"testing"
)

func TestProvisionServer(t *testing.T) {
	provision.ProvisionServer("example.com")
	expected := []byte(`server {
    server_name example.com;
    listen 80;

    error_log /var/log/example.com-nginx.log warn;

    
        location / {
    root /var/www/html;
    try_files $uri $uri.html $uri/;
}
    
}`)

	_, err := os.Stat("nginx-example.com")
	if err != nil {
		t.Error("File was not created")
	}
	f, err := ioutil.ReadFile("nginx-example.com")
	if err != nil {
		t.Error("Failed to read file")
	}
	if 	bytes.Equal(f, expected) {
		t.Errorf("File contained unexpect contents")
	}
}