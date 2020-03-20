package main

import (
	"encoding/json"
	"flag"
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
	flag.Parse()

	r := gin.Default()
	r.POST("/payload", func(c *gin.Context) {
		if c.Request.Body != nil {
			defer c.Request.Body.Close()
			var body []byte
			_, err := c.Request.Body.Read(body)
			log.Infof("%s", string(body))
			if err == nil {
				trigger := Trigger{}
				err = json.Unmarshal(body, trigger)
				if err == nil {
					log.Infof("%v", trigger)
				}
			}
		}
	})
	_ = r.Run(config.Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
