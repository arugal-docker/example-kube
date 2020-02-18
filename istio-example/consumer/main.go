package main

import (
	"fmt"
	"github.com/arugal-docker/example-kube/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	log   *logrus.Logger
	flags = struct {
		Addr     string
		Provider string
	}{
		Provider: "provider",
	}

	rootCmd = &cobra.Command{
		Use:          "consumer",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := os.Getenv("PROVIDER_HOSTNAME")
			if provider == "" {
				log.Printf("provider use flages: %s", flags.Provider)
				provider = flags.Provider
			}

			log.Printf("provider hostname:%s\n", provider)

			router := gin.Default()

			router.GET("/v1", func(ctx *gin.Context) {
				resp, err := request(fmt.Sprintf("http://%s:%s/v1", provider, "8081"))
				if err != nil {
					ctx.String(http.StatusInternalServerError, err.Error())
				} else {
					ctx.String(http.StatusOK, resp)
				}
			})
			return http.ListenAndServe(flags.Addr, router)
		},
	}
)

func init() {
	log = logger.Log
	rootCmd.PersistentFlags().StringVar(&flags.Addr, "addr", ":8080", "project-B addr")
	rootCmd.PersistentFlags().StringVar(&flags.Provider, "provider", "provider", "provider addr")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func request(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("request provider error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
