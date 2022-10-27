package main

import (
	selfLogger "GinHttps/logger"
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

var (
	Logger = selfLogger.InitLogger("HttpHttps")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	err := GinHttps()
	if err != nil {
		log.Fatal(err)
	}
}

func GinHttps() error {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(200, "{'message':'ok'}")
		Logger.Infof("Host=%s, URL=%s, Method=%s, Response=200", c.Request.Host, c.Request.URL, c.Request.Method)
	})

	go r.Run(":80")
	r.Use(TlsHandler(443))
	err := r.RunTLS(":443", "cert.pem", "key.pem")
	return err
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
