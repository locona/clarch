package clarch

import (
	"fmt"
	"html/template"
	"os"

	"github.com/locona/clarch/bindata"
)

type CmdInit struct{}

func (this *CmdInit) Init() {
	dirPath := "project"
	maps := make([]map[string]string, 0)
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/handler.tpl",
		"out": fmt.Sprintf("%s/handler.go", dirPath),
	})
	maps = append(maps, map[string]string{
		"tpl": "clarch/templates/project/repository.tpl",
		"out": fmt.Sprintf("%s/repository.go", dirPath),
	})

	for idx := range maps {
		if err := mkdir(dirPath); err != nil {
			panic(err)
		}

		tplFile, err := bindata.Asset(maps[idx]["tpl"])
		t := template.New("project")
		t, _ = t.Parse(string(tplFile))
		fp, err := os.OpenFile(maps[idx]["out"], os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		value := struct{}{}

		if err := t.Execute(fp, value); err != nil {
			panic(err)
		}
	}
}

func NewCmdInit() *CmdInit {
	return &CmdInit{}
}
