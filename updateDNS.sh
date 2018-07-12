#!/bin/sh
zid=""
domain=""
key=""
mail=""

did="" # domain id

if [ -z "$did" ]; then
    url="https://api.cloudflare.com/client/v4/zones/"$zid"/dns_records?type=A&name"$domain"&page=1&per_page=20&order=type&direction=desc&match=all"
    curl --insecure -X GET $url \
    -H "X-Auth-Email: $mail" \
    -H "X-Auth-Key: $key" \
    -H "Content-Type: application/json"
fi


ip=`curl -s http://api.ipify.org`
echo $ip

url="https://api.cloudflare.com/client/v4/zones/"$zid"/dns_records/"$did
curl --insecure -X PUT $url \
-H "X-Auth-Email: $mail" \
-H "X-Auth-Key: $key" \
-H "Content-Type: application/json" \
--data '{"type":"A","name":"'$domain'","content":"'$ip'","ttl":120,"proxied":false}'

