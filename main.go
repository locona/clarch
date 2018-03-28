package main

import (
	"flag"
	"os"
	"strings"

	"github.com/locona/clarch/clarch"
)

func main() {
	conf := parseArgs(os.Args...)

	dir, _ := os.Getwd()
	dirNames := strings.Split(dir, "/")
	conf.CurrentDir = dir
	conf.GithubUser = dirNames[len(dirNames)-2]
	conf.GithubRepository = dirNames[len(dirNames)-1]

	c := clarch.Clarch{Config: conf}
	c.Run()
}

func parseArgs(args ...string) clarch.Config {
	var conf clarch.Config
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.StringVar(&conf.Pkg, "pkg", "", "name or matching regular expression of interface to generate mock for")
	f.StringVar(&conf.DB, "db", "psql", "name or matching regular expression of interface to generate mock for")
	f.Parse(args[1:])
	return conf
}
