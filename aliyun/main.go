package main

import (
	"encoding/json"
	"github.com/arugal-docker/example-kube/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logger.Log
}

func main() {
	config := Config{Addr: ":8080"}
	config.addFlags()

	r := gin.Default()
	r.POST("/payload", func(c *gin.Context) {
		if c.Request.Body != nil {
			var body []byte
			n, err := c.Request.Body.Read(body)
			if err == nil && n > 0 {
				trigger := Trigger{}
				err = json.Unmarshal(body, trigger)
				if err == nil {
					log.Infof("%v", trigger)
				}
			}
		}
		c.Status(200)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(config.Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
