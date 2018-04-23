package clarch

import (
	"os"
)

type _new struct{}

func NewProject(project string) {
	if err := os.Mkdir(project, os.ModePerm); err != nil {
		panic(err)
	}

	if err := os.Chdir(project); err != nil {
		panic(err)
	}

	NewInit().Init()
}
