package main

import "flag"

type Config struct {
	Addr string
}

func (c *Config) addFlags() {
	flag.StringVar(&c.Addr, "addr", c.Addr, "server address")
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
