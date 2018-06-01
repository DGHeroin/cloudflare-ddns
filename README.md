./cloudflare-ddns -d domain -u mail -k cloudflare-key -z zones-id -i [new-IP|option] -x [domain id|option]
-d domain          域名
-u mail            CF用户名
-k cloudflare key  CF访问KEY
-z zones id        域名的zone id
-i new IP          新指定的IP(可选参数, 未指定则自动去查询)
-x domain id       子域名id(可选参数, 未指定则自动去CF查询)

## all by command
`
#!/bin/bash
# get domain id
curl --insecure -X GET "https://api.cloudflare.com/client/v4/zones/xx-zones_id-xx/dns_records?type=A&name=xx-your_domain-xx&page=1&per_page=20&order=type&direction=desc&match=all" \
-H "X-Auth-Email: your-mail@gmail.com" \
-H "X-Auth-Key: your-cf-key" \
-H "Content-Type: application/json" 

# change dns record
curl -X PUT "https://api.cloudflare.com/client/v4/zones/xx--zones_id-xx/dns_records/xx-domain_id-xx" \
     -H "X-Auth-Email: your-mail@gmail.com" \
     -H "X-Auth-Key: your-cf-key" \
     -H "Content-Type: application/json" \
     --data '{"type":"A","name":"xx-your-domain-xx","content":"xx--new_ip-xx","ttl":120,"proxied":false}'

`
