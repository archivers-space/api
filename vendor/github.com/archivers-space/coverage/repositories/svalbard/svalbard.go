package svalbard

import (
	"bufio"
	"github.com/archivers-space/archive"
	"github.com/archivers-space/coverage/tree"
	"net/url"
	"os"
	"strings"
)

var Repository = &repository{
	Id:          "1a1112f4-38a8-26d3-be70-c324bf686a87",
	Title:       "Project Svalbard",
	Description: "",
	Url:         "",
}

type repository archive.DataRepo

func (s *repository) GetId() string {
	return s.Id
}

func (r *repository) DataRepo() *archive.DataRepo {
	dr := archive.DataRepo(*r)
	return &dr
}

func (s *repository) AddUrls(t *tree.Node, src *archive.Source) error {
	f, err := os.Open("repositories/svalbard/svalbard_urls.txt")
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		node := t
		u, err := url.Parse(sc.Text())
		if err != nil {
			return err
		}
		if u.Scheme == "" {
			u.Scheme = "http"
		}

		// skip this url if it doesn't match the passed in Source filter
		if src != nil && !src.MatchesUrl(u.String()) {
			continue
		}

		node = node.Child(u.Scheme).Child(u.Host)
		components := strings.Split(u.Path, "/")

		for _, c := range components {
			if c != "" {
				node = node.Child(c)
			}
		}
		if u.RawQuery != "" {
			node = node.Child(u.RawQuery)
		}
	}
	return nil
}

func (s *repository) AddCoverage(t *tree.Node) {
	// no-op, need to add svalbard coverage work
}
