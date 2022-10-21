package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

var isHttps = flag.Bool("isHttps", true, "Use HTTPS")
var port int

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	err := GinHttps(*isHttps) // 这里false 表示 http 服务，非 https
	if err != nil {
		log.Fatal(err)
	}
}

func GinHttps(isHttps bool) error {
	if isHttps {
		port = 443
	} else {
		port = 80
	}

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(200, "{'message':'ok'}")
	})

	if isHttps {
		r.Use(TlsHandler(port))

		return r.RunTLS(":"+strconv.Itoa(port), "cert.pem", "key.pem")
	}

	return r.Run(":" + strconv.Itoa(port))
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
