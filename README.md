# GolangGinHttps 
* **建http server listen 80 or https server listen 443**
* **執行時以參數isHttps決定**
* **所有url都返回ok**



## compile for linux
```shell
GOOS=linux GOARCH=amd64 go build -o https-80 .
```

```shell
sudo nohup ./thomas-server-http-https -isHttps=true > /dev/null 2>&1 &
sudo nohup ./thomas-server-http-https -isHttps=false > /dev/null 2>&1 &
```