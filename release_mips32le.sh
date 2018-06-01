mkdir -p release
export GOROOT=/opt/mipsgo
export PATH=/opt/mipsgo/bin:$PATH

CGO_ENABLED=0 GOOS=linux GOARCH=mips32le go build -o release/cloudflare-ddns_linux_mips32le cloudflare-ddns.go
