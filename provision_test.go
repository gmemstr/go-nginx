package provision

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestProvisionServer(t *testing.T) {
	locationPlain := CreateLocation("/plain", false, "", false)
	locationProxy := CreateLocation("/proxy", true, "http://127.0.0.1:4343", false)
	locationWebsocket := CreateLocation("/ws", true, "http://127.0.0.1:4343", true)
	server := CreateServer("example.com", []string{locationPlain, locationProxy, locationWebsocket})
	err := ioutil.WriteFile("test-config", server.Bytes(), 0644)
	if err != nil {
		t.Error("Unable to write out file.")
	}

	expected := []byte(`server {
    server_name example.com;
    listen 80;

    error_log /var/log/example.com-nginx.log warn;

    
        location /plain {
    root /var/www/html/plain;
    try_files $uri $uri.html $uri/;
}
    
        location /proxy {
    proxy_pass http://127.0.0.1:4343;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
	proxy_http_version 1.1;
	
}
    
        location /ws {
    proxy_pass http://127.0.0.1:4343;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
	proxy_http_version 1.1;
	
	proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "Upgrade";
	
}
    
}
`)

	f, err := ioutil.ReadFile("test-config")
	if err != nil {
		t.Error("File was not created or is not readable.")
	}
	if !bytes.Equal(f, expected) {
		t.Error("Expected and actual config vary.")
	}
	err = os.Remove("test-config")
	if err != nil {
		t.Error("Unable to remove test file, please do so manually.")
	}
}