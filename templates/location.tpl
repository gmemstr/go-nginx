location {{ .Location }} {
    root {{ .Root }};
    try_files $uri $uri.html $uri/;
}