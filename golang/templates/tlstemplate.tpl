{{ with secret "ORGTLSCA/issue/ROLE" "common_name=CNAME" "ttl=TTL" "alt_names=localhost,CNAME" "ip_sans=127.0.0.1"}}
{{ .Data.CERT }}
{{ end }}