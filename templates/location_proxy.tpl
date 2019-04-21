location {{ .Location }} {
    proxy_pass {{ .ProxyAddress }};
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
	proxy_http_version 1.1;
	{{ if .IsWebsocket }}
	proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "Upgrade";
	{{ end }}
}