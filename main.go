package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	GinHttps(true) // 这里false 表示 http 服务，非 https
}

func GinHttps(isHttps bool) error {

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(200, "{'message':'ok'}")
	})

	if isHttps {
		r.Use(TlsHandler(80))

		return r.RunTLS(":"+strconv.Itoa(80), "cert.pem", "key.pem")
	}

	return r.Run(":" + strconv.Itoa(80))
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
