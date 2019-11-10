package main

import "github.com/trungtvq/craw/cmd"

var revision string

func main() {
	cmd.SetRevision(revision)
	cmd.Execute()
}
