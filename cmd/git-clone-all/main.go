package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	xfmt "github.com/fdschonborn/x/fmt"
	xexec "github.com/fdschonborn/x/os/exec"
	"github.com/spf13/pflag"
)

type Repos []Repo

type Repo struct {
	FullName string `json:"full_name"`
	CloneUrl string `json:"clone_url"`
	SshUrl   string `json:"ssh_url"`
}

const (
	defaultBaseUrl = "https://api.github.com"
	reposEndpoint  = "users/%s/repos"
)

var (
	baseUrl = pflag.StringP("base-url", "b", defaultBaseUrl, "Base API Url")
	user    = pflag.StringP("user", "u", "", "User to clone repositories from")
	method  = pflag.StringP("method", "m", "https", "Clone method to use (https or ssh)")
)

func main() {
	if err := run(); err != nil {
		xfmt.Eprintfln("Error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	pflag.Parse()

	if *user == "" {
		return errors.New("no user given")
	}

	resp, err := http.Get(*baseUrl + "/" + fmt.Sprintf(reposEndpoint, *user))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var repos Repos
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return err
	}

	for _, repo := range repos {
		var url string
		switch *method {
		case "https":
			url = repo.CloneUrl
		case "ssh":
			url = repo.SshUrl
		default:
			return fmt.Errorf("invalid method option: %s", *method)
		}

		gitcmd := xexec.NewBuilder("git").
			Args([]string{"clone", url}).
			Stderr(os.Stderr).
			Stdout(os.Stdout).
			Stdin(os.Stdin).
			Build()
		if err := gitcmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
