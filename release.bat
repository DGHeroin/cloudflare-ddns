@md release

@SET CGO_ENABLED=0
@SET GOOS=windows
@SET GOARCH=amd64
go build -o release/cloudflare-ddns_win.exe cloudflare-ddns.go

@SET CGO_ENABLED=0
@SET GOOS=darwin
@SET GOARCH=amd64
go build -o release/cloudflare-ddns_darwin cloudflare-ddns.go

@SET CGO_ENABLED=0
@SET GOOS=linux
@SET GOARCH=amd64
go build -o release/cloudflare-ddns_linux cloudflare-ddns.go

@SET CGO_ENABLED=0
@SET GOOS=linux
@SET GOARCH=mips
go build -o release/cloudflare-ddns_linux_mips cloudflare-ddns.go 

@SET CGO_ENABLED=0
@SET GOOS=linux
@SET GOARCH=mipsle
go build -o release/cloudflare-ddns_linux_mipsle cloudflare-ddns.go 

@SET CGO_ENABLED=0
@SET GOOS=linux
@SET GOARCH=mips32le
go build -o release/cloudflare-ddns_linux_mips32le cloudflare-ddns.go 