tee curl-format.txt <<- 'EOF'
time_namelookup: %{time_namelookup}\n
time_connect: %{time_connect}\n
time_appconnect: %{time_appconnect}\n
time_redirect: %{time_redirect}\n
time_pretransfer: %{time_pretransfer}\n
time_starttransfer: %{time_starttransfer}\n
----------\n
time_total: %{time_total}\n\n
EOF

DOMAIN="bing.com"
curl -w '@curl-format.txt' $DOMAIN
