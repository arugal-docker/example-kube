package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type Config struct {
	Addr          string
	kubeConfig    string
	Namespace     string
	GotifyAddress string
	GotifyToken   string
}

func (c *Config) addFlags() {
	flag.StringVar(&c.Addr, "addr", c.Addr, "server address")
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&c.kubeConfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&c.kubeConfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.StringVar(&c.Namespace, "namespace", c.Namespace, "kubernetes namespace")
	flag.StringVar(&c.GotifyAddress, "gotify-addr", c.GotifyAddress, "gotify address")
	flag.StringVar(&c.GotifyToken, "gotify-token", c.GotifyToken, "gotify token")
}

type Trigger struct {
	PushData   PushData   `json:"push_data"`
	Repository Repository `json:"repository"`
}

type PushData struct {
	Digest   string `json:"digest"`
	PushedAt string `json:"pushed_at"`
	Tag      string `json:"tag"`
}

type Repository struct {
	DateCreated            string `json:"date_created"`
	Name                   string `json:"name"`
	Namespace              string `json:"namespace"`
	Region                 string `json:"region"`
	RepoAuthenticationType string `json:"repo_authentication_type"`
	RepoFullName           string `json:"repo_full_name"`
	RepoOriginType         string `json:"repo_origin_type"`
	RepoType               string `json:"repo_type"`
}
