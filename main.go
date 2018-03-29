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
	switch conf.Mode {
	case "init":
		c.Init()
	default:
		c.Run()
	}
}

func parseArgs(args ...string) clarch.Config {
	var conf clarch.Config
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.StringVar(&conf.Mode, "mode", "run", "init | run")
	f.StringVar(&conf.Pkg, "pkg", "", "")
	f.StringVar(&conf.DB, "db", "gorm", "")
	f.Parse(args[1:])
	return conf
}
