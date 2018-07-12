#!/bin/sh
zid=""
domain=""
key=""
mail=""
#did="" #option

#get domain id
GetDomainID() {
    tmp=/tmp/cloudflare-ddns.xxxx
    url="https://api.cloudflare.com/client/v4/zones/"$zid"/dns_records?type=A&name"$domain"&page=1&per_page=20&order=type&direction=desc&match=all"
    curl -s --insecure -X GET $url \
    -H "X-Auth-Email: $mail" \
    -H "X-Auth-Key: $key" \
    -H "Content-Type: application/json" | sed 's/}},/\n/g'|grep $domain |sed 's/,/\n/g' | grep \"id\" | sed 's/"/ /g' | awk '{print $4}' > $tmp
    did=$(cat /tmp/cloudflare-ddns.xxxx)
    rm $tmp
}

#update domain IP
UpdateDNS() {
    ip=`curl -s http://api.ipify.org`
    url="https://api.cloudflare.com/client/v4/zones/"$zid"/dns_records/"$did
    curl --insecure -X PUT $url \
    -H "X-Auth-Email: $mail" \
    -H "X-Auth-Key: $key" \
    -H "Content-Type: application/json" \
    --data '{"type":"A","name":"'$domain'","content":"'$ip'","ttl":120,"proxied":false}'
}

if [ -z "$did" ]; then
    GetDomainID
fi

UpdateDNS
