package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
	"flag"
)

func die(msg string)  {
	fmt.Println(msg)
	os.Exit(0)
}

func GetMyIP() string {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	return string(ip)
}

func GetDomainId(domain string, zones_id string, mail string, key string) string{
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records?type=A&name=%s&page=1&per_page=20&order=type&direction=desc&match=all",
		zones_id, domain)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Email", mail)
	req.Header.Set("X-Auth-Key", key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		die(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// 解析数据
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)

	var ok bool // sucess
	err = json.Unmarshal(*objmap["success"], &ok)

	if ok == false {
		fmt.Println(string(body))
		return ""
	}

	var result []json.RawMessage // 解析出reslut
	err = json.Unmarshal(*objmap["result"], &result)

	var result_obj map[string]*json.RawMessage // 解析result array
	err = json.Unmarshal(result[0], &result_obj)

	var id string // 解析id
	err = json.Unmarshal(*result_obj["id"], &id)
	return id
}

func UpdateDNS(domain string, ip string, zones_id string, domain_id string, mail string, key string) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", zones_id, domain_id)
	post_body := fmt.Sprintf("{\"type\":\"A\",\"name\":\"%s\",\"content\":\"%s\",\"ttl\":120,\"proxied\":false}", domain, ip)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(post_body)))
	req.Header.Set("X-Auth-Email", mail)
	req.Header.Set("X-Auth-Key", key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		die(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ParseConfig() (string, string, string, string, string, string) {
	domain := flag.String("d", "", "domain")
	new_ip := flag.String("i", "", "IP")
	mail := flag.String("u", "", "email")
	key := flag.String("k", "", "key")
	zones_id := flag.String("z", "", "zones id")
	domain_id := flag.String("x", "", "domain id")

	if *new_ip == "" {
		*new_ip = GetMyIP()
	}

	if *domain_id == "" {
		*domain_id = ""
	}

	flag.Parse()
	if *domain == "" || *new_ip == "" || *mail == "" || *key == "" || *zones_id == "" {
		die("args error:\n" +
		 	"\ndomain:"+ *domain +
			"\nIP:"+ *new_ip +
			"\nmail:"+ *mail +
			"\nkey :"+ *key +
			"\nzones_id:"+ *zones_id +
			"\ndomain_id:"+ *domain_id)
	}
	return *domain, *new_ip, *mail, *key, *zones_id, *domain_id
}

func main(){
	domain, new_ip, mail, key, zones_id, domain_id := ParseConfig()
	if domain_id == "" {
		domain_id = GetDomainId(domain, zones_id, mail, key)
	}
	if (domain_id == "") {
		die("domain_id not found")
	}
	UpdateDNS(domain, new_ip, zones_id, domain_id, mail, key)
}
