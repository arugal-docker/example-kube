package main

import (
	"flag"
	"github.com/arugal-docker/example-kube/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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
	r.POST("/trigger", func(c *gin.Context) {
		if c.Request.Body != nil {
			defer c.Request.Body.Close()
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Errorf("Read body err: %v", err)
				return
			}
			log.Infof("%s", string(body))
		}
	})
	_ = r.Run(config.Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
