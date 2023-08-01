package model

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	gitURLs "github.com/whilp/git-urls"
)

type Repo struct {
	URL *url.URL
}

func NewRepo(path string) (*Repo, error) {
	u, err := gitURLs.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize repo from %s, due to: %w", path, err)
	}

	if u.Scheme != "ssh" {
		return nil, fmt.Errorf("%s is an unsupported scheme for repo URLs, only ssh is supported", u.Scheme)
	}

	return &Repo{URL: u}, nil
}

func (r *Repo) Name() string {
	ext := filepath.Ext(r.URL.Path)

	name := strings.TrimSuffix(r.URL.Path, ext)

	return strings.TrimSpace(name)
}

func (r *Repo) NameWithoutOrg() string {
	nameComponents := strings.Split(r.Name(), "/")

	return nameComponents[len(nameComponents)-1]
}
