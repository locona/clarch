package clarch

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/locona/clarch/bindata"
)

type handler struct {
	http http
}

type http struct {
	cmdadd
}

type uc struct {
	cmdadd
}
type repo struct {
	cmdadd
}

type model struct {
	cmdadd
}

type value struct {
	CurrentDir  string
	CurrentRepo string
	CurrentUser string
	CamelPkg    string
	Pkg         string
}

type cmdadd struct {
	value value
}

func gen(maps []map[string]string, value value) {
	for idx := range maps {
		fp, err := openOrCreate(maps[idx]["out"])
		defer fp.Close()
		if err != nil {
			panic(err)
		}
		b, err := bindata.Asset(maps[idx]["tpl"])
		if err != nil {
			panic(err)
		}

		t := template.New("gen")
		t, _ = t.Parse(string(b))
		t = t.Funcs(template.FuncMap{
			"safeHTML": func(s interface{}) template.HTML {
				return template.HTML(fmt.Sprint(s))
			},
		})

		if err := t.Execute(fp, value); err != nil {
			panic(err)
		}
	}
}

func (this *http) mapTplAndOutPath() []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/handler/http/http.tpl",
		"out": fmt.Sprintf("%s/handler/http/%s_handler.go", this.value.Pkg, this.value.Pkg),
	})

	m = append(m, map[string]string{
		"tpl": "clarch/templates/handler/http/params.tpl",
		"out": fmt.Sprintf("%s/handler/http/params.go", this.value.Pkg),
	})

	return m
}

func (this *uc) mapTplAndOutPath() []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/usecase/usecase.tpl",
		"out": fmt.Sprintf("%s/usecase/%s_uc.go", this.value.Pkg, this.value.Pkg),
	})
	return m
}

func (this *repo) mapTplAndOutPath() []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/repository/repository.tpl",
		"out": fmt.Sprintf("%s/repository/repository.go", this.value.Pkg),
	})
	return m
}

func (this *model) mapTplAndOutPath() []map[string]string {
	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"tpl": "clarch/templates/model.tpl",
		"out": fmt.Sprintf("model/%s.go", this.value.Pkg),
	})
	return m
}

func (this *cmdadd) Add() {
	GenHttp(this)
	GenUC(this)
	GenRepo(this)
	GenModel(this)
}

func NewCmdAdd(pkg string) *cmdadd {
	dir, _ := os.Getwd()
	dirNames := strings.Split(dir, "/")
	return &cmdadd{
		value: value{
			CurrentDir:  dir,
			CurrentUser: dirNames[len(dirNames)-2],
			CurrentRepo: dirNames[len(dirNames)-1],
			CamelPkg:    camelCase(pkg),
			Pkg:         pkg,
		},
	}
}

func GenHttp(cmd *cmdadd) {
	o := http{*cmd}
	gen(o.mapTplAndOutPath(), cmd.value)
}

func GenUC(cmd *cmdadd) {
	o := uc{*cmd}
	gen(o.mapTplAndOutPath(), cmd.value)
}

func GenRepo(cmd *cmdadd) {
	o := repo{*cmd}
	gen(o.mapTplAndOutPath(), cmd.value)
}

func GenModel(cmd *cmdadd) {
	o := model{*cmd}
	gen(o.mapTplAndOutPath(), cmd.value)
}
