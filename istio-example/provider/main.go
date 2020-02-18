package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	flags = struct {
		Addr     string
		HostName string
	}{}

	rootCmd = &cobra.Command{
		Use:          "provider",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			router := gin.Default()

			router.GET("/v1", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, fmt.Sprintf("hello, world! from %s", flags.HostName))
			})
			return http.ListenAndServe(flags.Addr, router)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.Addr, "addr", ":8081", "provider addr")
	flags.HostName, _ = os.Hostname()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
