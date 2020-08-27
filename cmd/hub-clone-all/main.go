package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	xexec "github.com/fdschonborn/x/os/exec"
	"github.com/spf13/pflag"
	"golang.org/x/crypto/ssh/terminal"
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
	baseUrl     = pflag.StringP("base-url", "b", defaultBaseUrl, "Base API Url")
	user        = pflag.StringP("user", "u", "", "User to clone repositories from")
	method      = pflag.StringP("method", "m", "https", "Clone method to use (https or ssh)")
	private     = pflag.BoolP("private", "p", false, "Clone private repositories (asks for password)")
	accessToken = pflag.StringP("access-token", "t", "", "Personal access token (will be asked otherwise)")
)

func main() {
	log.SetPrefix("hub-clone-all: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	pflag.Parse()

	if *user == "" {
		return errors.New("no user")
	}

	req, err := http.NewRequest(http.MethodGet, *baseUrl+"/"+fmt.Sprintf(reposEndpoint, *user), nil)
	if err != nil {
		return err
	}

	if *private {
		if *accessToken == "" {
			fmt.Print("Access Token: ")
			token, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}

			*accessToken = string(token)
		}

		req.SetBasicAuth(*user, *accessToken)
	}

	resp, err := http.DefaultClient.Do(req)
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

		err := xexec.NewBuilder("git", "clone", url).
			Stderr(os.Stderr).
			Stdout(os.Stdout).
			Stdin(os.Stdin).
			Build().
			Run()
		if err != nil {
			return err
		}
	}

	return nil
}
