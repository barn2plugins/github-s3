package main

import (
	"os"

	githubs3 "github.com/barn2plugins/github-s3"
)

func main() {
	githubs3.Run(
		func() string {
			return os.Getenv("GITHUB_SESSION")
		},
	)
}
