package pulls

import (
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

var (
	path = "./pullipfs/"
)

// Basic example of how to clone a repository using clone options.
func Pullrepo(url string) {

	// Clone the given repository to the given directory
	Info("git clone " + url)

	git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

}
