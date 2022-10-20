# GolangGinHttps

## compile for linux
```shell
GOOS=linux GOARCH=amd64 go build -o https-80 .
```

```shell
sudo nohup ./thomsd-server-http-https -isHttps=true > /dev/null 2>&1 &
sudo nohup ./thomsd-server-http-https -isHttps=false > /dev/null 2>&1 &
```