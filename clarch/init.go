package clarch

import (
	"fmt"
	"html/template"
	"os"

	"github.com/locona/clarch/bindata"
)

type CmdInit struct{}

func (this *CmdInit) Init() error {
	dirPath := "project"
	if err := mkdir(dirPath); err != nil {
		return err
	}

	tplFile, err := bindata.Asset("clarch/templates/project/handler.tpl")
	t := template.New("project")
	t, _ = t.Parse(string(tplFile))
	fp, err := os.OpenFile(fmt.Sprintf("%s/handler.go", dirPath), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	value := struct{}{}
	return t.Execute(fp, value)
}

func NewCmdInit() *CmdInit {
	return &CmdInit{}
}
