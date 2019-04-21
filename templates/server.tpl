server {
    server_name {{ .ServerName }};
    listen 80;

    error_log {{ .LogLocation }} warn;

    {{ range .Locations }}
        {{ . }}
    {{ end }}
}
