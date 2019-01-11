package gitstar

import (
	"github.com/google/go-github/github"
)

// GitStar defines main struct of the app
type GitStar struct {
	client *github.Client
}

// New provides create of the main app
func New(client *github.Client) *GitStar {
	return &GitStar{
		client: client,
	}
}

// Client returns github client
func (g *GitStar) Client() *github.Client {
	return g.client
}
